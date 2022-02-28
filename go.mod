module example

go 1.17

require gee v0.0.0

replace gee => ./gee

//多个mod文件可能会导致项目莫名其妙的引用失败
