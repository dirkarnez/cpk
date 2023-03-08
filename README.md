cpk
===
**Work in progress**  

A real stupid C/C++ package manager for own studying (intended not to use other C++ package managers available).  
**And YES! `cpk` sounds exactly like swearing in Cantonese-mixing-English. hahaha :P**

### Convention / Rules / Targets
1. `cpk` focuses on providing **prebuilt libraries** for CMake projects
2. `cpk` must be portable (no registry, no system directory access, no `%PATH%` modifying, no `pkg-config`, no crazy path-searching)
    - store the fetched C++ packages in its directory instead of %USERPROFILE% or `~`
3. `cpk` must support different compilers
    - For Windows, mingw and MSVC are a must. 
4. `TARGET_COMPILER` should be well-defined and consistent, with compiler version
5. `TARGET_TAG` should stick to the original versioning practice (eg. some may version "v2.5.1", some may version "2.5.1")
6. `TARGET_LIBRARY` is the library name, it should follow the original name
7. There should be a centralized repository store the information of prebuilt libraries (exactly like `npm`)
8. The creation of prebuilt libraries and their repository should be automated in GitHub action and repository template (`gcc -dumpversion` gets "8.1.0")
9. `cpk` and the supporting library repositories should also support listing include paths for easier development in VSCode
10. For C/C++ projects using `cpk`, there should not be any configuration files needed - `cpk` should only depend on (read / write) explicit `cmake` file complying project onfiguration (`find_package` in `CMakeLists.txt`)
11. `cpk` should leave room for modifying third-party libraries (for example: editing original `CMakeLists.txt` and may / may not submit PR in a **fork repo** and rebuilding prebuilts, including ad-hoc /private redirection of libraries
13. `cpk` should resolve dependency-related issues like nested dependency
14. `cpk` should give hints (using CMake macros?) on further setup for the projects using the installed libraries
15. `cpk` should be unobtrusive, play well with unsupported / local libraries - `cpk` should have a dedicated folder for its CMake modules fetching and configuration (use `CMAKE_INSTALL_PREFIX`?)
16. `cpk` should have a logic to redirect to official git repo (for header-only libraries like [g-truc/glm: OpenGL Mathematics (GLM)](https://github.com/g-truc/glm))
17. **`cpk` should be available as a CMake module**
    - refer to [dirkarnez/cmake-playground](https://github.com/dirkarnez/cmake-playground)
    
### How to use `cpk` in C/C++ project
1. Prerequisites for C/C++ setup (also these compiler, git, cmake, `cpk` are in `%PATH%`)
2. Having `cmake/modules/.gitkeep` - `cpk` assumes `cmake/cpk_modules` exists (or warn) and expect to overwrite the FindXXX.cmake if applicable
3. `cpk` assumes there are version constraints (exact or ranged, **`master` should not be supported**)

### Supporting libraries
- [ ] [SFML](https://github.com/dirkarnez/sfml-prebuilt), WIP
	- `v2.5.1`
- [ ] [alaingalvan/CrossWindow](https://github.com/dirkarnez/crosswindow-prebuilt), WIP
	- `c3511e43d0bbdc5fd23b4717580cc47ad3df87e4c`
- [ ] [alaingalvan/CrossWindow-Graphics](https://github.com/dirkarnez/crosswindow-graphics-prebuilt), WIP
	- `c28022e3ba565a975e8bf6dada1aa605c6f73a848`
- [ ] [cycfi/elements](https://github.com/dirkarnez/elements-prebuilt), WIP
	- `cd7106afe85fe99c28a3d7314a7ea5d3bd50e9c84`
- [ ] [dacap/clip](https://github.com/dirkarnez/clip-prebuilt), WIP
	- `v1.4`
- [ ] [GiovanniDicanio/WinReg](https://github.com/dirkarnez/WinReg-prebuilt), WIP
	- `v4.1.2`
- [ ] [sciplot/sciplot](https://github.com/dirkarnez/sciplot-prebuilt), WIP
	- `v0.2.2`
- [ ] [boostorg/boost](https://github.com/dirkarnez/boost-prebuilt), WIP
	- `v1.80.0`
- [ ] [arvidn/libtorrent](https://github.com/dirkarnez/libtorrent-prebuilt), WIP
	- `v2.0.7`
- [ ] [dpilger26/NumCpp](https://github.com/dirkarnez/NumCpp-prebuilt), WIP
	- `v2.8.0`
- [ ] [openbabel/openbabel](https://github.com/dirkarnez/openbabel-prebuilt), WIP
	- `v3.1.1`
- [ ] [facebook/zstd](https://github.com/dirkarnez/zstd-prebuilt), WIP
	- `v1.5.2`
- [ ] [ros/catkin](https://github.com/dirkarnez/catkin-prebuilt), WIP
    - `v0.8.10`
- [ ] [Dav1dde/glad](https://github.com/dirkarnez/glad2-prebuilt), WIP
    - `v2.0.4` 
- [ ] [lief-project/LIEF](https://github.com/dirkarnez/LIEF-prebuilt), WIP
	- `v0.12.3`
- [ ] [xianyi/OpenBLAS](https://github.com/dirkarnez/openblas-prebuilt), WIP
	- `v0.3.21`
- [ ] [knik0/faac](https://github.com/dirkarnez/faac-prebuilt), WIP
	- `v1.30`
- [ ] [zherczeg/sljit](https://github.com/dirkarnez/sljit-prebuilt)', WIP
    - `cfaecd2b12f3e1119019255ca547f7db5c8627a7d`
- [ ] [zsiciarz/aquila](https://github.com/dirkarnez/aquila-prebuilt), WIP
    - `cd5e3bde3c8d07678a95f225b657a7bb23d0c4730`
- [ ] [zxing-cpp/zxing-cpp](https://github.com/dirkarnez/zxing-cpp-prebuilt), WIP
    - `v1.4.0`
- [ ] [libuv/libuv](https://github.com/dirkarnez/libuv-prebuilt), WIP
	- `v1.44.2`
- [ ] [XZ Utils | SourceForge.net](https://github.com/dirkarnez/xz-prebuilt), WIP
	- `v5.4.0`
- [ ] [erincatto/box2d](https://github.com/dirkarnez/box2d-prebuilt), WIP
	- `v2.4.1`
- [ ] [coin-or/CppAD](https://github.com/dirkarnez/cppad-prebuilt), WIP
    - `stable/20210000`
- [ ] [maxbachmann/rapidfuzz-cpp](https://github.com/dirkarnez/rapidfuzz-cpp-prebuilt), WIP
	- `v1.10.1`
- [ ] [soundtouch/soundtouch](https://github.com/dirkarnez/soundtouch-prebuilt), WIP
	- `v2.3.1`
- [ ] [ImageMagick/ImageMagick: 🧙‍♂️ ImageMagick 7](https://github.com/ImageMagick/ImageMagick)
- [ ] [sreiter/ProMesh: Grid generation for scientific computations](https://github.com/sreiter/ProMesh)
- [ ] [socketio/socket.io-client-cpp](https://github.com/dirkarnez/socket.io-client-cpp-prebuilt), WIP
	- `v3.1.0`
- [ ] [CGAL/cgal](https://github.com/dirkarnez/cgal-prebuilt), WIP
	- `v5.5.1`
- [ ] [Voro++ - A 3D Voronoi cell software library](https://github.com/dirkarnez/voropp-prebuilt), WIP
	- `v0.4.6`
- [ ] [LAME](https://github.com/dirkarnez/lame-prebuilt), WIP
    - `v3.100.0`
- [ ] [json-c/json-c](https://github.com/dirkarnez/json-c-prebuilt), WIP
	- `v0.16`
- [ ] [Independent JPEG Group](https://github.com/dirkarnez/libjpeg-prebuilt), WIP
	- `v9e`
- [ ] [uchardet/uchardet](https://github.com/dirkarnez/uchardet-prebuilt)
    - `v0.0.8`
- [ ] [dlbeer/quirc](https://github.com/dirkarnez/quirc-prebuilt), WIP
    - `v1.1`
- [ ] [lipnitsk/libcue](https://github.com/dirkarnez/libcue-prebuilt), WIP
    - `v2.2.1`
- [ ] [jfalcou/kumi: C++20 Compact Tuple Tools](https://github.com/dirkarnez/kumi-prebuilt), WIP
    - `v3.0`
- [ ] [sreiter/stl_reader](https://github.com/dirkarnez/stl_reader-prebuilt), WIP
    - `v1.0.0`
- [ ] [fmtlib/fmt](https://github.com/dirkarnez/fmt-prebuilt), WIP
    - `v9.1.0`
- [ ] [arximboldi/lager](https://github.com/dirkarnez/lager-prebuilt), WIP
    - `c05e4ffda3eb1a4158eb678f21fbfc00e96a8c95d`
- [ ] [msteinbeck/tinyspline](https://github.com/dirkarnez/tinyspline-prebuilt), WIP
	- `v0.5.0`
- [ ] [audionamix/wave: Simple C++ library for reading and writting wave files](https://github.com/audionamix/wave)
- [ ] [UG4/ugcore: The core functionality of UG4. Includes sources, build-scripts, and utility scripts.](https://github.com/UG4/ugcore)
- [ ] [abseil/abseil-cpp: Abseil Common Libraries (C++)](https://github.com/abseil/abseil-cpp)
- [ ] [sbellus/libmicrohttpd-cmake: The repository contains cmake build files for libmicrohttpd (https://www.gnu.org/software/libmicrohttpd/)](https://github.com/sbellus/libmicrohttpd-cmake)
- [ ] [audionamix/auconv: An audio file conversion library](https://github.com/audionamix/auconv)
- [ ] [audionamix/rtff: A real time frequential filtering library](https://github.com/audionamix/rtff)
- [ ] [j256/dmalloc: Debug Malloc memory allocation debugging C library](https://github.com/j256/dmalloc)
- [ ] [microsoft/wil](https://github.com/dirkarnez/wil-prebuilt), WIP
	- `v1.0.220914.1`
- [ ] [PortMidi/portmidi](https://github.com/dirkarnez/portmidi-prebuilt), WIP
    - `v2.0.4`
- [ ] [PortAudio/portaudio](https://github.com/dirkarnez/portaudio-prebuilt), WIP
    - `v19.7.0`
- [ ] [FITSIO Home Page](https://github.com/dirkarnez/cfitsio-prebuilt), WIP
	- `v4.2.0`
- [ ] [gflags/gflags: The gflags package contains a C++ library that implements commandline flags processing. It includes built-in support for standard types such as string and the ability to define flags in the source file in which they are used. Online documentation available at:](https://github.com/gflags/gflags)
- [ ] [gpg/libgpg-error](https://github.com/dirkarnez/libgpg-error-prebuilt), WIP
	- `v1.46`
- [ ] [cycfi/q](https://github.com/dirkarnez/q-prebuilt)
- [ ] [facebookincubator/MY_ENUM: Small c++ macro library to add compile-time introspection to c++ enum classes.](https://github.com/facebookincubator/MY_ENUM)
- [ ] [yanyiwu/cppjieba](https://github.com/dirkarnez/cppjieba-prebuilt), WIP
    - `v5.0.3`
- [ ] [wxWidgets/wxWidgets](https://github.com/dirkarnez/wxWidgets-prebuilt), WIP
    - `v3.2.1`
- [ ] [AshampooSystems/boden: Purely native C++ cross-platform GUI framework for Android and iOS development. https://www.boden.io](https://github.com/AshampooSystems/boden/)
- [ ] [libsndfile/libsndfile: A C library for reading and writing sound files containing sampled audio data.](https://github.com/libsndfile/libsndfile), WIP
    - [dirkarnez/libsndfile-prebuilt](https://github.com/dirkarnez/libsndfile-prebuilt)
- [ ] [gnuradio/gnuradio: GNU Radio – the Free and Open Software Radio Ecosystem](https://github.com/gnuradio/gnuradio)
- [ ] [g-truc/glm: OpenGL Mathematics (GLM)](https://github.com/dirkarnez/glm-prebuilt), WIP
	- `cfc8f4bb442b9540969f2f3f351c4960d91bca17a`, WIP
- [ ] [pradeep-pyro/tinynurbs: C++ library for NURBS curves and surfaces](https://github.com/dirkarnez/tinynurbs-prebuilt), WIP
	- `c98ab6e55ffa441d7471f4ef39f520eaf712b2d74`
- [ ] [videolan/x264](https://github.com/dirkarnez/x264-prebuilt), WIP
    - `cbaee400fa9ced6f5481a728138fed6e867b0ff7f`
- [ ] [Download Xpdf and XpdfReader](https://www.xpdfreader.com/download.html)
- [ ] [ampl/mp](https://github.com/dirkarnez/mp-prebuilt), WIP
    - ~~`v3.1.0`~~
- [ ] [ampl/asl](https://github.com/dirkarnez/asl-prebuilt/tree/main), WIP
    - `c642206393ec127906fb248343e0bcee1c2119850`
- [ ] [gpac/gpac: Modular Multimedia framework for packaging, streaming and playing your favorite content.](https://github.com/gpac/gpac)
- [ ] [openframeworks/openFrameworks: openFrameworks is a community-developed cross platform toolkit for creative coding in C++.](https://github.com/openframeworks/openFrameworks)
- [ ] [ampl/ampls-api: AMPL Solver APIs](https://github.com/ampl/ampls-api)
- [ ] [libusb/libusb: A cross-platform library to access USB devices](https://github.com/libusb/libusb)
- [ ] [libusb/hidapi: A Simple library for communicating with HID devices on Linux, Mac and Windows.](https://github.com/libusb/hidapi)
- [ ] [twitter/vireo](https://github.com/dirkarnez/vireo-prebuilt), WIP
    - `c7f0bc1ad50f5eaf30a09de209f0172f2def36abd`
- [ ] [oberhumer.com: LZO real-time data compression library], WIP
    - `v2.10`
- [ ] [libexpat/libexpat](https://github.com/dirkarnez/libexpat-prebuilt), WIP
    - `v2.4.9`
- [ ] [libffi/libffi](https://github.com/dirkarnez/libffi-prebuilt), WIP
    - `v3.4.3`
- [ ] [enzo1982/mp4v2](https://github.com/dirkarnez/mp4v2-prebuilt), WIP
    - `v2.1.1`
- [ ] [pocoproject/poco](https://github.com/dirkarnez/poco-prebuilt), WIP
    - `v1.12.2`
- [ ] [GNU Readline](https://github.com/dirkarnez/readline-prebuilt), WIP
    - `v8.2`
- [ ] [coreutils/gnulib](https://github.com/dirkarnez/gnulib-prebuilt), WIP
    - `v0.1`
- [ ] [jugit.fz-juelich.de/mlz/libcerf](https://github.com/dirkarnez/libcerf-prebuilt), WIP
    - `v2.1`
- [ ] [IfcOpenShell/IfcOpenshell](https://github.com/dirkarnez/IfcOpenShell-prebuilt), WIP
    - `v0.7.0`
- [ ] [flann-lib/flann](https://github.com/dirkarnez/flann-prebuilt), WIP
    - `v1.9.1`
- [ ] [dlfcn-win32/dlfcn-win32](https://github.com/dirkarnez/dlfcn-win32-prebuilt), WIP
    - `v1.3.1`
- [ ] [gpg/gpg4win: GnuPG for Windows. NOTE: Maintainers are not tracking this mirror. Do not make pull requests here, nor comment any commits, submit them usual way to bug tracker (https://www.gnupg.org/documentation/bts.html) or to the mailing list (https://www.gnupg.org/documentation/mailing-lists.html).](https://github.com/gpg/gpg4win)
- [ ] [lfreist/hwinfo](https://github.com/dirkarnez/hwinfo-prebuilt), WIP
- [ ] [dirkarnez/libass-prebuilt](https://github.com/dirkarnez/libass-prebuilt), WIP
- [ ] [dirkarnez/libjit-prebuilt](https://github.com/dirkarnez/libjit-prebuilt), WIP
- [ ] [dirkarnez/bison-prebuilt](https://github.com/dirkarnez/bison-prebuilt), WIP
- [ ] [mpdecimal](https://github.com/dirkarnez/mpdecimal-prebuilt), WIP
- [ ] [GNU Guile](https://github.com/dirkarnez/guile-prebuilt), WIP
    - `v3.0.8`
- [ ] [xnd-project/gumath: Library supporting function dispatch on general data containers. C base and Python wrapper](https://github.com/xnd-project/gumath)
- [ ] [nemtrif/utfcpp](https://github.com/dirkarnez/utfcpp-prebuilt), WIP
    - `v3.2.1`
- [ ] [erincatto/box2d: Box2D is a 2D physics engine for games](https://github.com/erincatto/box2d)
- [ ] [assimp/assimp: The official Open-Asset-Importer-Library Repository. Loads 40+ 3D-file-formats into one unified and clean data structure.](https://github.com/assimp/assimp)
- [ ] [bulletphysics/bullet3: Bullet Physics SDK: real-time collision detection and multi-physics simulation for VR, games, visual effects, robotics, machine learning etc.](https://github.com/bulletphysics/bullet3)
- [ ] [mapbox/geojson-cpp: A C++14 library for converting GeoJSON into geometry.hpp representation](https://github.com/mapbox/geojson-cpp)
- [ ] [mapbox/tippecanoe: Build vector tilesets from large collections of GeoJSON features.](https://github.com/mapbox/tippecanoe)
- [ ] [wawanbreton/cvcomposer: Advanded GUI for OpenCV, to compose various filters and quickly see the result](https://github.com/wawanbreton/cvcomposer)
    - `ce30ab337ed625e1b868202b9e92b85bd9e3f766a`
- [ ] [bloomberg/bde: Basic Development Environment - a set of foundational C++ libraries used at Bloomberg.](https://github.com/bloomberg/bde)
- [ ] [facebook/watchman: Watches files and records, or triggers actions, when they change.](https://github.com/facebook/watchman)
- [ ] [facebook/yoga: Yoga is a cross-platform layout engine which implements Flexbox. Follow https://twitter.com/yogalayout for updates.](https://github.com/facebook/yoga)
- [ ] [facebook/fboss: Facebook Open Switching System Software for controlling network switches.](https://github.com/facebook/fboss)
- [ ] [facebook/wangle: Wangle is a framework providing a set of common client/server abstractions for building services in a consistent, modular, and composable way.](https://github.com/facebook/wangle)
- [ ] [fox-toolkit.org](https://github.com/dirkarnez/fox-prebuilt), WIP
- [ ] [nlohmann/json](https://github.com/dirkarnez/json-prebuilt), WIP
    - `v3.11.2`
- [ ] [rttrorg/rttr](https://github.com/dirkarnez/rttr-prebuilt), WIP
    - `v0.9.6`
- [ ] [Artelnics/opennn: OpenNN - Open Neural Networks Library](https://github.com/Artelnics/OpenNN)
- [ ] [pdfium - Git at Google](https://pdfium.googlesource.com/pdfium/)
- [ ] [g-truc/glm: OpenGL Mathematics (GLM)](https://github.com/g-truc/glm), WIP
    - `v0.9.9`
- [ ] [OpenVPN/openvpn: OpenVPN is an open source VPN daemon](https://github.com/OpenVPN/openvpn)
- [ ] [Kitware/VTK](https://github.com/dirkarnez/VTK-prebuilt), WIP
    - `v9.2.2`
- [ ] [MetaLogic Inference - Summary [Savannah]](https://savannah.gnu.org/projects/metalogic-inference)
- [ ] [v8/v8: The official mirror of the V8 Git repository](https://github.com/v8/v8)
- [ ] [openMVG/openMVG: open Multiple View Geometry library. Basis for 3D computer vision and Structure from Motion.](https://github.com/openMVG/openMVG)
- [ ] [facebook/kuduraft: A Raft Library in C++ based on the Raft implementation in Apache Kudu](https://github.com/facebook/kuduraft)
- [ ] [facebook/fatal: Fatal is a library for fast prototyping software in modern C++. It provides facilities to enhance the expressive power of C++. The library is heavily based on template meta-programming, while keeping the complexity under-the-hood.](https://github.com/facebook/fatal)
- [ ] [facebook/folly: An open-source C++ library developed and used at Facebook.](https://github.com/facebook/folly)
- [ ] [unicode-org/icu](https://github.com/dirkarnez/icu4c-prebuilt), WIP
    - `v71.1`
- [ ] [libharu/libharu: libharu - free PDF library](https://github.com/libharu/libharu), WIP
- [ ] [l-smash/l-smash](https://github.com/dirkarnez/l-smash-prebuilt), WIP
    - `v2.14.5`
- [ ] [dealii/dealii: The development repository for the deal.II finite element library.](https://github.com/dealii/dealii)
- [ ] [google/skia: Skia is a complete 2D graphic library for drawing Text, Geometries, and Images.](https://github.com/google/skia)
    - https://github.com/LibCMaker/LibCMaker_Skia
- [ ] [apache/xerces-c](https://github.com/dirkarnez/xerces-c-prebuilt), WIP
    - `v3.2.3`
- [ ] [LibreDWG/libredwg](https://github.com/dirkarnez/libredwg-prebuilt), WIP
    - `ce32962da3f247937f6355f80a113ea716016e998`
- [ ] [JagPDF - Downloads](http://www.jagpdf.org/downloads.htm)
    - [MAGANER/PdfBuilder: markdown-like language to generate pdf documents.](https://github.com/MAGANER/PdfBuilder)
- [ ] Qt
- [ ] [google/double-conversion](https://github.com/dirkarnez/double-conversion-prebuilt)
- [ ] [verilator/verilator: Verilator open-source SystemVerilog simulator and lint system](https://github.com/verilator/verilator)
- [ ] [llvm-project/pstl at main · llvm/llvm-project](https://github.com/llvm/llvm-project/tree/main/pstl)
- [ ] [google/fhir)(https://github.com/dirkarnez/fhir-prebuilt), WIP
- [ ] [armadillo](https://github.com/dirkarnez/armadillo-prebuilt), WIP
    - `v11.1.1`
- [ ] [**alecthomas/entityx**](https://github.com/alecthomas/entityx)
- [ ] [jackaudio/jack2: jack2 codebase](https://github.com/jackaudio/jack2)
- [ ] [cartographer-project/cartographer: Cartographer is a system that provides real-time simultaneous localization and mapping (SLAM) in 2D and 3D across multiple platforms and sensor configurations.](https://github.com/cartographer-project/cartographer)
- [ ] [alsa-project/alsa-lib: The Advanced Linux Sound Architecture (ALSA) - library](https://github.com/alsa-project/alsa-lib)
- [ ] [catchorg/Catch2](https://github.com/dirkarnez/catch2-prebuilt), WIP
    - `v3.1.0`
- [ ] OpenAL
- [ ] http://www.gphoto.org/proj/libgphoto2/
- [ ] https://c.msgpack.org/cpp/index.html
- [ ] [fltk/fltk](https://github.com/dirkarnez/fltk-prebuilt), WIP
    - `v1.3.8`
- [ ] [craigsapp/midifile](https://github.com/dirkarnez/midifile-prebuilt), WIP
    - `cf7ded879139061c41631797e5b735d6d5ceee23e`
- [ ] [FreeGLUTProject/freeglut](https://github.com/dirkarnez/freeglut-prebuilt), WIP
    - `v3.2.2`
- [ ] [FreeImage download | SourceForge.net](https://sourceforge.net/projects/freeimage/)
- [ ] https://glew.sourceforge.net/
- [ ] [taocpp/PEGTL](https://github.com/dirkarnez/pegtl-prebuilt), WIP
    - `v3.2.7`
- [ ] [libjpeg-turbo/libjpeg-turbo](https://github.com/dirkarnez/libjpeg-turbo-prebuilt), WIP
    - `v2.1.3`
- [ ] CImg
- [ ] [bblanchon/ArduinoJson](https://github.com/dirkarnez/arduinojson-prebuilt), WIP
    - `v6.19.4`
- [ ] [chemfiles/chemfiles](https://github.com/dirkarnez/chemfiles-prebuilt), WIP
    - `v0.10.2`
- [ ] [antirez/sds](https://github.com/dirkarnez/sds-prebuilt), **non-cmake** repo, **WIP**
- [ ] [kuafuwang/LspCpp](https://github.com/dirkarnez/lspcpp-prebuilt), WIP
    - `c6e280bccc1df85527089abcd8a37098b0d6c1eb9`
    - [dirkarnez/lspcpp-boilerplate](https://github.com/dirkarnez/lspcpp-boilerplate)
- [ ] [fuzzylite/fuzzylite](https://github.com/dirkarnez/fuzzylite-prebuilt), **WIP**
    - `v6.0`
    - [dirkarnez/fuzzylite-boilerplate](https://github.com/dirkarnez/fuzzylite-boilerplate)
- [ ] https://github.com/mlpack/mlpack
- [ ] [glfw/glfw](https://github.com/dirkarnez/glfw-prebuilt), WIP
	- `v3.3.8`
- [ ] [thestk/rtaudio: A set of C++ classes that provide a common API for realtime audio input/output across Linux (native ALSA, JACK, PulseAudio and OSS), Macintosh OS X (CoreAudio and JACK), and Windows (DirectSound, ASIO, and WASAPI) operating systems.](https://github.com/thestk/rtaudio)
- [ ] [OGRECave/ogre-next](https://github.com/OGRECave/ogre-next)
- [ ] [lballabio/QuantLib](https://github.com/dirkarnez/quantlib-prebuilt), WIP
    - `v1.27`
- [ ] [symisc/unqlite](https://github.com/dirkarnez/unqlite-prebuilt), WIP
	- `v1.1.9`
- [ ] [sqlite/sqlite](https://github.com/dirkarnez/sqlite-prebuilt), WIP
    - `v3.40.1`
- [ ] [open-mpi/ompi: Open MPI main development repository](https://github.com/open-mpi/ompi)
- [ ] [OpenMP/sources: This repository contains supplementary source code for the OpenMP(R) API Specification.](https://github.com/OpenMP/sources)
- [ ] [oneapi-src/oneTBB: oneAPI Threading Building Blocks (oneTBB)](https://github.com/oneapi-src/oneTBB)
- [ ] [oneapi-src/oneDNN: oneAPI Deep Neural Network Library (oneDNN)](https://github.com/oneapi-src/oneDNN)
- [ ] [oneapi-src/oneCCL: oneAPI Collective Communications Library (oneCCL)](https://github.com/oneapi-src/oneCCL)
- [ ] [oneapi-src/oneDAL: oneAPI Data Analytics Library (oneDAL)](https://github.com/oneapi-src/oneDAL)
- [ ] [oneapi-src/oneMKL: oneAPI Math Kernel Library (oneMKL) Interfaces](https://github.com/oneapi-src/oneMKL)
- [ ] [oneapi-src/oneDPL: oneAPI DPC++ Library (oneDPL) https://software.intel.com/content/www/us/en/develop/tools/oneapi/components/dpc-library.html](https://github.com/oneapi-src/oneDPL)
- [ ] [laserpants/dotenv-cpp](https://github.com/dirkarnez/dotenv-cpp-prebuilt), WIP
    - `cca1d543f78702921dbb398c865806b71dedfd2eb`
- [ ] [Armadillo: C++ library for linear algebra & scientific computing](http://arma.sourceforge.net/)
- [ ] libjpeg
- [ ] [glennrp/libpng](https://github.com/dirkarnez/libpng-prebuilt), WIP
    - `v1.6.34`
- [ ] [qgis/QGIS: QGIS is a free, open source, cross platform (lin/win/mac) geographical information system (GIS)](https://github.com/qgis/QGIS)
- [ ] [graphql/libgraphqlparser: A GraphQL query parser in C++ with C and C++ APIs](https://github.com/graphql/libgraphqlparser)
- [ ] [opencv/opencv](https://github.com/dirkarnez/opencv-prebuilt), WIP
    - `ce14c3cff856ad0f56fb14a5e59f9bb9849555903`
- [ ] Magick++
- [ ] https://github.com/facebook/folly
- [ ] Simd
- [ ] https://www.openal.org/
- [ ] https://www.kfrlib.com/
- [ ] [zeux/pugixml](https://github.com/dirkarnez/pugixml-prebuilt), WIP
	- `v1.12.1`
- [ ] https://github.com/STEllAR-GROUP/hpx/
- [ ] [GNU gettext - Git Repositories [Savannah]](https://github.com/dirkarnez/gettext-prebuilt), WIP
- [ ] [GNU libiconv - Git Repositories [Savannah]](https://github.com/dirkarnez/libiconv-prebuilt), WIP
- [ ] [GNOME / libxml2 · GitLab](https://gitlab.gnome.org/GNOME/libxml2)
- [ ] [http://mpg123.org/](https://github.com/dirkarnez/mpg123-prebuilt), WIP
	- `c90e9d67613c506d599af806c81221c2e9dc77079`
- [ ] https://github.com/LRN/ntldd
- [ ] [LRN/libntlink: A library that implements symlink(), link(), unlink(), lstat() and readlink() on top of WinAPI](https://github.com/LRN/libntlink)
- [ ] [gpg/gpg4win](https://github.com/gpg/gpg4win)
- [ ] cpp-netlib
- [ ] [pybind/pybind11: Seamless operability between C++11 and Python](https://github.com/pybind/pybind11)
- [ ] https://github.com/snakster/cpp.react
- [ ] [hosseinmoein/DataFrame](https://github.com/dirkarnez/dataframe-prebuilt), WIP
    - `v1.21.0`
- [ ] [FFTW/fftw3](https://github.com/dirkarnez/fftw3-prebuilt), WIP
    - `v3.3.10`
- [ ] [oatpp/oatpp](https://github.com/dirkarnez/oatpp-prebuilt), WIP
    - `v1.3.0`
- [ ] [madler/zlib](https://github.com/dirkarnez/zlib-prebuilt), WIP
    - `v1.2.12`
- [ ] OpenMP
- [ ] https://github.com/steven-varga/h5cpp
- [ ] [GMP](https://github.com/dirkarnez/gmp-prebuilt), WIP
    - `v6.2.1`
- [ ] [GNU MPFR](https://github.com/dirkarnez/mpfr-prebuilt), WIP
    - `v4.1.0`
- [ ] [GNU MPC — multiprecision.org](https://github.com/dirkarnez/mpc-prebuilt), WIP
    - `v1.2.1`
- [ ] GStreamer
- [ ] [libsdl-org/SDL](https://github.com/dirkarnez/sdl-prebuilt), WIP
    - `v2.0.2`
- [ ] TagLib
- [ ] cpprestsdk
- [ ] [google/brotli](https://github.com/dirkarnez/brotli-prebuilt), WIP
    - `v1.0.9`
- [ ] Google tests
- [ ] gRPC
- [ ] [davisking/dlib](https://github.com/dirkarnez/dlib-prebuilt), WIP
    - `v19.24`
- [ ] darknet
- [ ] https://github.com/GHamrouni/Recommender
- [ ] [libeigen / eigen · GitLab](https://github.com/dirkarnez/eigen-prebuilt), WIP
    - `v3.4.0`
- [ ] https://github.com/softwareQinc/qpp
- [ ] https://github.com/Mbed-TLS/mbedtls
- [ ] https://github.com/libevent/libevent
    - https://github.com/dirkarnez/libevent-prebuilt
- [ ] https://github.com/the-tcpdump-group/libpcap
- [ ] [protocolbuffers/protobuf](https://github.com/dirkarnez/protobuf-prebuilt), WIP
    - `v21.7`
- [ ] https://github.com/USCiLab/cereal
- [ ] [FFmpeg/FFmpeg](https://github.com/dirkarnez/ffmpeg-prebuilt)
- [ ] https://github.com/jakogut/tinyvm
- [ ] CGAL
- [ ] CUDA
- [ ] [jbeder/yaml-cpp](https://github.com/dirkarnez/yaml-cpp-prebuilt), WIP
    - `v0.7.0`
- [ ] [The Programming Language Lua](https://github.com/dirkarnez/lua-prebuilt), WIP
    - `v5.4.4`
- [ ] [OpenBSD bcrypt](https://github.com/dirkarnez/bcrypt-prebuilt), WIP
    - ``
- [ ] https://github.com/c-ares/c-ares.git
- [ ] [curl/curl](https://github.com/dirkarnez/curl-prebuilt), WIP
    - `v7.86.0`
- [ ] https://github.com/curl/doh
- [ ] https://www.gnutls.org/
- [ ] https://github.com/rockdaboot/libpsl
- [ ] https://www.libssh2.org/
- [ ] [hosseinmoein/DataFrame](https://github.com/dirkarnez/dataframe-prebuilt), WIP
    - `v1.21.0`
- [ ] https://www.cgal.org/
- [ ] https://github.com/DieShell/TclTk
- [ ] https://github.com/dirkarnez/boost-prebuilt
- [ ] https://github.com/FreeRDP/FreeRDP
- [ ] [manticoresoftware/manticoresearch: Easy to use open source fast database for search | Good alternative to Elasticsearch now | Drop-in replacement for E in the ELK soon](https://github.com/manticoresoftware/manticoresearch/)
- [ ] [openssl/openssl](https://github.com/dirkarnez/openssl-prebuilt), WIP
    - `v3.0.5`
- [ ] [RaftLib/RaftLib: The RaftLib C++ library, streaming/dataflow concurrency via C++ iostream-like operators](https://github.com/RaftLib/RaftLib)
- [ ] [Xiph.Org / flac · GitLab](https://github.com/dirkarnez/flac-prebuilt), WIP
    - `v1.4.2`
- [ ] [Xiph.Org / Ogg · GitLab](https://gitlab.xiph.org/xiph/ogg)
- [ ] GNU Scientific Library
    - [ ] [ampl/gsl](https://github.com/dirkarnez/gsl-prebuilt), WIP
        - `v2.7.0`
- [ ] [freedesktop/fontconfig: Font customization and configuration library](https://github.com/freedesktop/fontconfig)
    - https://github.com/LibCMaker/LibCMaker_FontConfig
- [ ] [electronicarts/EASTL: EASTL stands for Electronic Arts Standard Template Library. It is an extensive and robust implementation that has an emphasis on high performance.](https://github.com/electronicarts/EASTL)
- [ ] [ETLCPP/etl: Embedded Template Library](https://github.com/ETLCPP/etl)
- [ ] [jfalcou/eve: Expressive Vector Engine - SIMD in C++ Goes Brrrr](https://github.com/jfalcou/eve)
- [ ] [g-truc/glm: OpenGL Mathematics (GLM)](https://github.com/g-truc/glm)
- [ ] [HDFGroup/hdf5](https://github.com/dirkarnez/hdf5-prebuilt), WIP
    - `v1.13.2`  
- [ ] [oneapi-src/oneTBB: oneAPI Threading Building Blocks (oneTBB)](https://github.com/oneapi-src/oneTBB)
- [ ] [modm-io/avr-libstdcpp: Subset of the C++ standard library for AVR targets](https://github.com/modm-io/avr-libstdcpp)
- [ ] [webview/webview: Tiny cross-platform webview library for C/C++/Golang. Uses WebKit (Gtk/Cocoa) and Edge (Windows)](https://github.com/webview/webview)
- [ ] [andrew-gresyk/HFSM2: High-Performance Hierarchical Finite State Machine Framework](https://github.com/andrew-gresyk/HFSM2)
- [ ] [agenium-scale/nsimd: Agenium Scale vectorization library for CPUs and GPUs](https://github.com/agenium-scale/nsimd/)
- [ ] [jfalcou/eve: Expressive Vector Engine - SIMD in C++ Goes Brrrr](https://github.com/jfalcou/eve)
- [ ] [dmlc/xgboost: Scalable, Portable and Distributed Gradient Boosting (GBDT, GBRT or GBM) Library, for Python, R, Java, Scala, C++ and more. Runs on single machine, Hadoop, Spark, Dask, Flink and DataFlow](https://github.com/dmlc/xgboost)
- [ ] [freedesktop/cairo: cairo's central repository](https://github.com/freedesktop/cairo)
- [ ] [hselasky/libumidi: MIDI file and interface library](https://github.com/hselasky/libumidi)
- [ ] [hselasky/midipp: MIDI Player Pro](https://github.com/hselasky/midipp)
- [ ] [hselasky/jack_umidi: Raw MIDI to JACK daemon](https://github.com/hselasky/jack_umidi)
- [ ] [bzip2/bzip2](https://github.com/dirkarnez/bzip2-prebuilt), WIP
    - `v1.0.8`
- [ ] [harfbuzz/harfbuzz](https://github.com/dirkarnez/harfbuzz-prebuilt), WIP
    - `v5.3.0`
- [ ] [freetype/freetype: Official mirror of https://gitlab.freedesktop.org/freetype/freetype](https://github.com/freetype/freetype)

### TODO
- [GNOME · GitLab](https://gitlab.gnome.org/GNOME)
- [freedesktop · Groups · Explore · GitLab](https://gitlab.freedesktop.org/explore/groups)
- [GnuWin32 Packages](https://gnuwin32.sourceforge.net/packages.html)
- [VideoLAN · GitLab](https://code.videolan.org/videolan)
- [VTK · GitLab](https://gitlab.kitware.com/vtk)
- [Kitware, Inc.](https://github.com/Kitware)
- [Xiph.Org · GitLab](https://gitlab.xiph.org/xiph)
- [ros2/ros2.repos at rolling · ros2/ros2](https://github.com/ros2/ros2/blob/rolling/ros2.repos)
- [QEMU · GitLab](https://gitlab.com/qemu-project/)
- [fleur / fleur · GitLab](https://iffgit.fz-juelich.de/fleur/fleur)
