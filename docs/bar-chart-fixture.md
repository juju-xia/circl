# Bar Chart Fixture

This branch contains deterministic, non-cryptographic changes for exercising a
repository comparison chart. The changes intentionally avoid production code so
they can be used for visual and numerical checks without affecting CIRCL's
security-sensitive packages.

## Dimensions To Check

| Dimension | What the chart should show |
| --- | --- |
| Files changed | The number of paths in `git diff --name-only` |
| Additions | The sum of the first column in `git diff --numstat` |
| Deletions | The sum of the second column in `git diff --numstat` |
| Net lines | Additions minus deletions |
| Commits | The number of commits in the selected revision range |

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
