fmon
====
tool for executing given command upon file change.

example:
  fmon -deadtime=1000 -path=monitorDir dir

run "dir" when files/folders under "monitorDir" changes.
It won't execute "dir" command until 1000ms has passed.

