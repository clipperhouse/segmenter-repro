An attempt to compare results of

- https://github.com/blevesearch/segment
- https://github.com/clipperhouse/uax29

Have a look at the test file.

## What I've found (June 2022)

- Bleve splits a run of spaces into separate tokens, while uax29 returns a single token with multiple spaces
- Bleve appears to be Unicode 8.0.0, uax29 is 13.0.0, seems like a difference on emoji skin tone modifiers
