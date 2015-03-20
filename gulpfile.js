/*jslint node: true */
'use strict';

var gulp = require('gulp');

// var nodemon = require('gulp-nodemon');
// var jest = require('gulp-jest');

var cssmin = require('gulp-cssmin');
var livereload = require('gulp-livereload');
// var env = require('gulp-env');


var browserifyTask = require('./gulp/browserify');
var styleTask = require('./gulp/style');
var goTask = require('./gulp/go');

var dest = './build',
  src = './src/client',
  mui = './node_modules/material-ui/src',
  dependencies = [
  'react',
  'react-addons',
  'fluxible',
  'react-tools',
];  

var config = {
  server:{
    main: '.'
  },
  Less: {
    src: src + '/styles/less/main.less',
    watch: [src +'/styles/**/*.less', src+'/app/**/*.less'],
    dest: dest
  },
  injectPaths:{
    src:'./gulp/paths.js',
    dest:'./'
  },
  fonts: {
    src: mui + '/less/material-design-fonticons/fonts/**',
    dest: dest + '/fonts/mdfonticon'
  },
  muiFonts: {
    src: mui + '/less/material-ui-icons/fonts/**',
    dest: dest + '/fonts'
  },
  Browserify: {
    // Enable source maps
    debug: true,
    // A separate bundle will be generated for each
    // bundle config in the list below
    bundleConfigs: [{
      entries: src + '/app/client.jsx',
      dest: dest,
      outputName: 'main.js',
      dependencies: dependencies
    }]
  }
};


gulp.task('serve', ['browserify','less','watch'], function () {
  // lr().listen(35729,function(){
  livereload.listen(35729, function(){
    console.log('live reload listening on 35729');
  });
});

gulp.task('go',function(){
  goTask(config.server,{
    development:true
  });
});

gulp.task('watch',function(){
  gulp.watch([config.Less.watch],['less']);
});

gulp.task('browserify',function(){
  browserifyTask(config.Browserify,{
    development: true
  });
});

gulp.task('less',function(){
  styleTask(config.Less,{
    development: true
  });
});


gulp.task('deploy', function () {

  browserifyTask(config.Browserify,{
    development: false
  });
  
  styleTask(config.Less,{
    development: false,
  });

});

// gulp.task('set-env', function () {
//     env({
//       file:'./server/config/local.env'
//     });
// });

// gulp.task('inject-paths',function(){
//   injectPathsTask(config.injectPaths);
// });

// gulp.task('nodemon', function(done){
//   nodemon({ script: './index', env: { 'NODE_ENV': 'development'}})
//   .on('restart');
// });

// The default task (called when you run `gulp` from cli)
gulp.task('default', ['serve']);


// gulp.task('jest', function () {
//   return gulp.src('./client/app/src/').pipe(jest({
//     scriptPreprocessor: "./preprocessor.js",
//     unmockedModulePathPatterns: [
//         "node_modules/react"
//     ],
//     testDirectoryName: "tests",
//     testPathIgnorePatterns: [
//         "node_modules",
//         "spec/support"
//     ],
//     moduleFileExtensions: [
//         "js",
//         "json",
//         "react"
//     ]
//   }));
// });

