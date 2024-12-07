package day05

import (
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"

	"github.com/kabourneeak/aoc2024/internal/days"
)

func Run(input string, w io.Writer) error {

	model, err := parseInput(input)
	if err != nil {
		return err
	}

	part1, err := part1(model)
	if err != nil {
		return err
	}

	part2, err := part2(model)
	if err != nil {
		return err
	}

	fmt.Fprintf(w, "Part 1 answer is %d\n", part1)
	fmt.Fprintf(w, "Part 2 answer is %d\n", part2)

	return nil
}

type inputModel struct {
	OrderingRules []orderingRule
	Updates       []update
}

type orderingRule struct {
	Pred int
	Succ int
}

type update struct {
	Pages []int
}

func parseInput(input string) (*inputModel, error) {

	lines := days.ToLines(input)

	parseOrderingRule := func(line string) (*orderingRule, error) {
		parts := strings.Split(line, "|")

		pred, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, err
		}

		succ, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}

		rule := &orderingRule{
			Pred: pred,
			Succ: succ,
		}

		return rule, nil
	}

	parseUpdate := func(line string) (*update, error) {
		parts := strings.Split(line, ",")

		pages := make([]int, 0, len(parts))

		for _, p := range parts {
			page, err := strconv.Atoi(p)
			if err != nil {
				return nil, err
			}

			pages = append(pages, page)
		}

		u := &update{
			Pages: pages,
		}

		return u, nil
	}

	rules := make([]orderingRule, 0)
	updates := make([]update, 0)

	for _, line := range lines {
		if line == "" {
			// skip this line
		} else if strings.Contains(line, "|") {
			rule, err := parseOrderingRule(line)
			if err != nil {
				return nil, err
			}

			rules = append(rules, *rule)
		} else {
			update, err := parseUpdate(line)
			if err != nil {
				return nil, err
			}

			updates = append(updates, *update)
		}
	}

	model := &inputModel{
		OrderingRules: rules,
		Updates:       updates,
	}

	return model, nil
}

func part1(input *inputModel) (int, error) {
	/*
		Find all rules that apply to the Update
		- Rules s.t. Pred and Succ both appear in the Update

		For each page P with value X;
		- filter the dependencies for that page; all rules where Succ == P
		- we know from above that Pred for each of these appears in the Update and so _is_
		  a valid dependency for P
		- verify all dependencies are present
	*/

	sum := 0

	for _, update := range input.Updates {
		rules := rulesForUpdate(input.OrderingRules, &update)

		if validateUpdate(&update, rules) {
			sum += scoreUpdate(&update)
		}
	}

	return sum, nil
}

func part2(input *inputModel) (int, error) {
	sum := 0

	/*
		Find all rules that apply to the Update
		- Rules s.t. Pred and Succ both appear in the Update

		See if the Update is already valid, and skip in that case

		To fix an update, for each page P with value X;
		- filter the dependencies for that page; all rules where Succ == P
		- if a dependency D is found which appears after P in the update order,
		  swap P and D; now D < P as required.
		- if all dependencies are met for P, then move to next page.
		- If P was swapped further down, then it will be visited again at some point
		  and if it still has an outstanding dependency, it will be moved down again.
	*/

	for _, update := range input.Updates {
		rules := rulesForUpdate(input.OrderingRules, &update)

		// don't score updates which are already valid
		if validateUpdate(&update, rules) {
			continue
		}

		fixedUpdate := fixUpdate(&update, rules)

		if !validateUpdate(fixedUpdate, rules) {
			return 0, fmt.Errorf("not valid after fixing:\noriginal %v\n   fixed %v", update.Pages, fixedUpdate.Pages)
		}

		sum += scoreUpdate(fixedUpdate)
	}

	return sum, nil
}

func rulesForUpdate(sourceRules []orderingRule, update *update) (out []orderingRule) {
	for _, rule := range sourceRules {
		if slices.Contains(update.Pages, rule.Pred) && slices.Contains(update.Pages, rule.Succ) {
			out = append(out, rule)
		}
	}

	return
}

func validatePage(page int, preds []int, rules []orderingRule) bool {
	deps := days.Filter(rules, func(r orderingRule) bool { return r.Succ == page })

	for _, dep := range deps {
		if !slices.Contains(preds, dep.Pred) {
			return false
		}
	}

	return true
}

func validateUpdate(update *update, rules []orderingRule) bool {
	for i, page := range update.Pages {
		if !validatePage(page, update.Pages[:i], rules) {
			return false
		}
	}

	return true
}

func fixPage(page_i int, pages []int, rules []orderingRule) (wasChanged bool) {
	page := pages[page_i]
	pageDeps := days.Filter(rules, func(r orderingRule) bool { return r.Succ == page })

	for _, dep := range pageDeps {
		pred_i := slices.Index(pages, dep.Pred)

		if pred_i > page_i {
			days.Swap(pages, page_i, pred_i)
			return true
		}
	}

	// no changes were made
	return false
}

func fixUpdate(in *update, rules []orderingRule) *update {

	pages := make([]int, len(in.Pages))
	copy(pages, in.Pages)

	for i := 0; i < len(pages); {

		if fixPage(i, pages, rules) {
			// if the current page was swapped, then a new page resides at i
			// so we loop again with the same i
			continue
		}

		// this page is good, move on to the next
		i += 1
	}

	return &update{Pages: pages}
}

func scoreUpdate(update *update) int {
	mid := len(update.Pages) / 2
	return update.Pages[mid]
}
