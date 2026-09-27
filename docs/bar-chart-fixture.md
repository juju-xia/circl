# Bar Chart Fixture

This branch contains deterministic, non-cryptographic changes for exercising a
repository comparison chart. The changes intentionally avoid production code and
use documentation plus test-only code, so they can be used for visual and
numerical checks without affecting CIRCL's security-sensitive packages.

## Dimensions To Check

| Dimension | What the chart should show |
| --- | --- |
| Files changed | The number of paths in `git diff --name-only` |
| Additions | The sum of the first column in `git diff --numstat` |
| Deletions | The sum of the second column in `git diff --numstat` |
| Net lines | Additions minus deletions |
| Total lines | Additions plus deletions |
| Commits | The number of commits in the selected revision range |
| Per category | The chart's grouping should match the selected file paths |

## Useful Revision Ranges

For the complete branch comparison:

```sh
git diff --numstat main...test/bar-chart
git log --oneline main..test/bar-chart
```

For one commit at a time:

```sh
git show --numstat --oneline <commit>
```

The fixture includes multiple small commits so a chart can be checked against
both per-commit values and the cumulative branch total.

## Numerical Checks

Use machine-readable output when verifying a chart:

```sh
git diff --numstat main...test/bar-chart
git diff --shortstat main...test/bar-chart
git rev-list --count main..test/bar-chart
```

Compare the plotted file count against `git diff --name-only main...test/bar-chart
| wc -l`. If a chart groups changes by directory or file type, verify the sum of
all groups before checking their relative bar lengths.
