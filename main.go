package main

import (
	"encoding/csv"
	"flag"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/icza/screp/rep"
	"github.com/icza/screp/repparser"
	"github.com/marianogappa/raszagal/analyzer"
)

func main() {
	var (
		_analyzers = map[string]analyzer.Analyzer{
			(&analyzer.IsThereAZerg{}).Name(): &analyzer.IsThereAZerg{},
			(&analyzer.MyAPM{}).Name():        &analyzer.MyAPM{},
			(&analyzer.MyRace{}).Name():       &analyzer.MyRace{},
		}
		flags      = map[string]*bool{}
		fReplay    = flag.String("replay", "", "path to replay file")
		fReplays   = flag.String("replays", "", "comma-separated paths to replay files")
		fReplayDir = flag.String("replay-dir", "", "path to folder with replays (recursive)")
		fMe        = flag.String("me", "", "comma-separated list of player names to identify as the main player")
	)
	for name, a := range _analyzers {
		flags[name] = flag.Bool(name, false, a.Description())
	}
	flag.Parse()
	var (
		analyzers     = map[string]analyzer.Analyzer{}
		csvFieldNames = []string{} // TODO add filename
	)
	for name, f := range flags {
		if *f {
			analyzers[name] = _analyzers[name]
			csvFieldNames = append(csvFieldNames, name)
		}
	}

	// Prepares for CSV output
	sort.Strings(csvFieldNames)
	w := csv.NewWriter(os.Stdout)
	w.Write(csvFieldNames)

	// Prepares AnalyzerContext
	ctx := analyzer.AnalyzerContext{Me: map[string]struct{}{}}
	if fMe != nil && len(*fMe) > 0 {
		for _, name := range strings.Split(*fMe, ",") {
			ctx.Me[strings.TrimSpace(name)] = struct{}{}
		}
	}

	// Parse replay filename flags
	var replays = map[string]struct{}{}
	*fReplay = strings.TrimSpace(*fReplay)
	if len(*fReplay) >= 5 && (*fReplay)[len(*fReplay)-4:] == ".rep" {
		replays[*fReplay] = struct{}{}
	}
	if *fReplays != "" {
		for _, r := range strings.Split(*fReplays, ",") {
			r = strings.TrimSpace(r)
			if len(r) >= 5 && r[len(r)-4:] == ".rep" {
				replays[r] = struct{}{}
			}
		}
	}
	if *fReplayDir != "" {
		e := filepath.Walk(*fReplayDir, func(path string, info os.FileInfo, err error) error {
			if err == nil && len(info.Name()) >= 5 && info.Name()[len(info.Name())-4:] == ".rep" {
				r := path
				replays[r] = struct{}{}
			}
			return nil
		})
		if e != nil {
			log.Fatal(e)
		}
	}

	// Main loop parsing replays
	for replay := range replays {
		analyzerInstances := make(map[string]analyzer.Analyzer, len(analyzers))
		for n, a := range analyzers {
			analyzerInstances[n] = a
		}

		r, err := repparser.ParseFile(replay)
		if err != nil {
			log.Printf("Failed to parse replay: %v\n", err)
			continue
		}
		tryCompute(r)

		var results = map[string]analyzer.Result{}
		for name, a := range analyzerInstances {
			if a.StartReadingReplay(r, ctx) {
				results[name], _ = a.IsDone()
				delete(analyzerInstances, name)
			}
		}
		for _, c := range r.Commands.Cmds {
			for name, a := range analyzerInstances {
				if a.ProcessCommand(c) {
					results[name], _ = a.IsDone()
					delete(analyzerInstances, name)
				}
			}

		}

		csvRow := make([]string, 0, len(csvFieldNames))
		for _, field := range csvFieldNames {
			csvRow = append(csvRow, results[field].Value())
		}
		w.Write(csvRow)
	}
	w.Flush()
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}

func tryCompute(r *rep.Replay) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered panic: %v", r)
		}
	}()
	r.Compute()
}