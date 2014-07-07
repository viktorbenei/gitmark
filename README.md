gitmark
=======

Command line git bookmarking


# TODO

GO: Minimal Git app
- through config files
- can add new repo (clone) -> a step-by-step guide, input from user, similar to SourceTree
- fetch all -> list repos with changes
- 'open'/'goto' -> cd into the dir


Gomake / gorelease
- go test
- go build
- switch to master, tag, switch back to dev


[BUGFIX] -> ~/.gitmarkrc is not ok, ~ is not expanded!
//- read config file (JSON)
//- list: list them
//- check: uncommited changes
//- scan: search for local git repo folders + option: save into gitmarkrc if not yet saved
//- open: open the repository with a specified command (os/exec)
- add: w/ a step-by-step guide, input from user, similar to SourceTree
- add v2: 2 modes, 1) add local working copy 2) clone a repo
- check: git fetch on all, checks if pull/push (w flag)
- status: more detailed info of all
- remove a gitmark, with option to remove the repository as well (delete folder)

- cleanup of tags (leave only the last X)
- get the next version-tag (increment the last one by one)

- alias: in addition to the 'title', a shorter 'shortcut' ID
- pull all branch: try to pull, every local branch

- scan -prune : remove repos not found during the scan
