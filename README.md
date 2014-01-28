[![Build Status](https://drone.io/github.com/tshm/fmon/status.png)](https://drone.io/github.com/tshm/fmon/latest)

fmon
====
tool for executing given command upon file change.

example:
  fmon -deadtime=1000 -path=monitorDir dir

run "dir" when files/folders under "monitorDir" changes.
It won't execute "dir" command until 1000ms has passed.

