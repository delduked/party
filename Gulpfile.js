const gulp = require('gulp');
const pug = require('gulp-pug');
const sass = require('gulp-sass')(require('sass'));
const autoprefixer = require('gulp-autoprefixer');
const browsersync = require('browser-sync').create();
const ts = require('gulp-typescript');
const tsProject = ts.createProject('tsconfig.json');

// BrowserSync for live reload on file saves
function browserSync(done) {
  browsersync.init({
    server: {
      baseDir: ['dist'],
      directory: false,
      stream: true
    },
    // Specifiy a socket implementation of your website
    //socket: {

       // For using a socket implementation of your website,
       // configure your DNS records to point to the ip address of yourhost machine.
       // Specify the socket name next to the field 'namespace' equalto your chosen DNS record pointing
       // to you host machine. ex: 'your.dns.record'

       // namespace: "gulp.mintymint.info"
    //},
    // disable scroll, click, page refresh etc... to be disabled across all clients
    // ghostMode: true,
    port: 8080
  });
  done();
}
// BrowserSync Reload
function browserSyncReload(done) {
  browsersync.reload();
  done();
}
// BrowserSync Reload

// pug Preprocessor
function pugPro() {
   return (
      gulp.src('src/pug/index.pug')
      // Specifies which file will be processed into html
      .pipe(pug({
          pretty: true
      }))
      // Compiles the pug file into HTML
      .pipe(gulp.dest('dist'))
      // Specifies where the processed HTML file will reside
      .pipe(browsersync.reload({stream: true}))
   );
};
// pug Preprocessor

// Sass Preprocessor
function sassPro() {
   return (
      gulp.src('src/sass/style.sass')
      .pipe(sass()) //converts sass to css
      .pipe(autoprefixer('last 2 version', 'safari 5', 'ie 8', 'ie 9','ff 17', 'opera 12.1', 'ios 6', 'android 4'))
      .pipe(gulp.dest('dist'))
      .pipe(browsersync.reload({stream: true}))
   );
};
// Sass Preprocessor

// TypeScript Preprocessor
function tsPro(){
   var tsResult = gulp.src('src/ts/**/*.ts') // or tsProject.src()
   .pipe(tsProject());

   return tsResult.js.pipe(gulp.dest('dist'));
}
// TypeScript Preprocessor

// Watch files
function watchFiles() {

   //gulp.watch(['src/css/**/*.sass','src/pug/**/*.pug','src/ts/**/*.ts'], gulp.series([sassPro,pugPro,tsPro]));

   gulp.watch('src/sass/**/*.sass',sassPro);
   gulp.watch('src/pug/**/*.pug',pugPro);
   gulp.watch('src/ts/**/*.ts',tsPro);
   
   gulp.watch(['dist/index.html','dist/style.css','dist/index.js'], gulp.series(browserSyncReload));
};
// Watch files

// define tasks to process
const build = gulp.series(
   pugPro,
   sassPro, 
   tsPro,
   browserSync, 
   watchFiles
);
// define tasks to process

// export tasks
exports.default = build;
