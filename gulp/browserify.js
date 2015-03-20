/* browserify task
   ---------------
   Bundle javascripty things with browserify!
   This task is set up to generate multiple separate bundles, from
   different sources, and to use Watchify when run from the default task.
   See browserify.bundleConfigs in gulp/config.js
*/

var browserify   = require('browserify');
var reactify = require('reactify'); 
var watchify     = require('watchify');
var bundleLogger = require('./utils/bundleLogger');
var gulp         = require('gulp');
var handleErrors = require('./utils/handleErrors');
var source       = require('vinyl-source-stream');
var livereload = require('gulp-livereload');
var gulpif = require('gulp-if');
var uglify = require('gulp-uglify');
var streamify = require('gulp-streamify');

var browserifier =function(config, options, callback) {

  var bundleQueue = config.bundleConfigs.length;

  var browserifyThis = function(bundleConfig) {

    var bundler = browserify({
      // Required watchify args
      cache: {}, packageCache: {}, fullPaths: false,
      // Specify the entry point of your app
      entries: bundleConfig.entries,
      // Add file extentions to make optional in your requires
      extensions: config.extensions,
      transform: [reactify], // We want to convert JSX to normal javascript
      // Enable source maps!
      debug: config.debug
    });
    // We set our dependencies as externals on our app bundler when developing    
    

    var bundle = function() {
      // Log when bundling starts
      bundleLogger.start(bundleConfig.outputName);
      return bundler
        .bundle()
        // Report compile errors
        .on('error', handleErrors)
        // Use vinyl-source-stream to make the
        // stream gulp compatible. Specifiy the
        // desired output filename here.
        .pipe(source(bundleConfig.outputName))
        .pipe(gulpif(!options.development, streamify(uglify())))
        // Specify the output destination
        .pipe(gulp.dest(bundleConfig.dest))
        .pipe(gulpif(options.development,livereload()))
        .on('end', reportFinished);
    };

    if(options.development) {
      bundleLogger.announce('watching');
      // Wrap with watchify and rebundle on changes
      watcher = watchify(bundler);
      // Rebundle on update

      watcher.on('update', bundle);
    }

    if(options.development){
      bundleConfig.dependencies.forEach(function (dep) {
        bundler.external(dep);
      });

      var vendorsBundler = browserify({
        debug: true,
        require: bundleConfig.dependencies,
        transform: [reactify] // We want to convert JSX to normal javascript
      });
      
      bundleLogger.start('vendors.js');
      // Run the vendor bundle
      var start = new Date();
      vendorsBundler.bundle()
        .on('error', handleErrors)
        .pipe(source('vendors.js'))
        .pipe(gulpif(!options.development, streamify(uglify())))
        .pipe(gulp.dest(bundleConfig.dest))
        .on('end', function(){
          bundleLogger.end('vendors.js');
        });
    }



    function reportFinished() {
      // Log when bundling completes
      bundleLogger.end(bundleConfig.outputName);

      if(bundleQueue) {
        bundleQueue--;
        if(bundleQueue === 0) {
          // If queue is empty, tell gulp the task is complete.
          // https://github.com/gulpjs/gulp/blob/master/docs/API.md#accept-a-callback
          if (typeof callback === 'function'){
            callback();
          }
        }
      }
    }

    return bundle();
  };

  // Start bundling with Browserify for each bundleConfig specified
  config.bundleConfigs.forEach(browserifyThis);
};

module.exports=browserifier;
