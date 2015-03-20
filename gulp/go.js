// gulp     = require 'gulp'
// symlink  = require 'gulp-sym'
// shell    = require 'gulp-shell'
// settings = require './settings'
// path     = require 'path'

// gulp.task 'go', () ->
//   gulp
//     .src(settings.server.root)
//     .pipe(symlink("#{settings.server.gohome}/src/#{settings.server.pkg}", { force: true }))
//     .pipe(shell("GOPATH=$GOPATH:`pwd` go build -o ../#{settings.server.binary} #{settings.server.pkg}/main", { cwd: settings.server.gohome }))

/* jlint node:true*/
var gulp = require('gulp');
var symlink = require('gulp-sym');
var shell = require('gulp-shell');
var path = require('path');

function go(config,options,callback){
  gulp.src(config.main)
  // .pipe(symlink())
  .pipe(shell("go build -o ./build/main ./src/server/main"));
}

module.exports = go;