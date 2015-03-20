/*jshint node:true*/
'use strict';
var gulp = require('gulp'),
  less = require('gulp-less'),
  autoprefixer = require('gulp-autoprefixer'),
  sourcemaps = require('gulp-sourcemaps'),
  handleErrors = require('./utils/handleErrors'),
  gulpif = require('gulp-if'),
  cssmin = require('gulp-cssmin'),
  livereload = require('gulp-livereload');
  var inject = require('gulp-inject');

module.exports = function(config, options){
  return gulp.src(config.src)
    .pipe(sourcemaps.init())
    .pipe(inject(
      gulp.src([config.src+'/app/**/*.less'], {read: false}), {
        starttag: '// injector:less',
        endtag: '// endinjector',
        transform: function (filePath) {
          if (filePath.slice(-5) === '.less') {
            filePath = filePath.replace('/client/', '../../');
            console.log(filePath);
            return '@import \'' + filePath + '\';';
          }
          return inject.transform.apply(inject.transform, arguments);
        }
      }
    ))
    .pipe(less())
    .on('error', handleErrors)
    .pipe(gulpif(!options.development,cssmin()))
    .pipe(autoprefixer({cascade: false, browsers: ['last 2 versions']}))
    .pipe(sourcemaps.write())
    .pipe(gulpif(options.development,livereload()))
    .pipe(gulp.dest(config.dest));
};