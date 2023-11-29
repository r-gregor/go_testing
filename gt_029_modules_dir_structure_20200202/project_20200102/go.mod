module github.com/r-gregor/project_20200102

go 1.13

require github.com/r-gregor/tmspkg v1.0.0

require github.com/r-gregor/mymodule v1.0.0

replace github.com/r-gregor/mymodule => ./mymodule

replace github.com/r-gregor/tmspkg => ./tmspkg
