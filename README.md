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

### How to use `cpk` in C/C++ project
1. Prerequisites for C/C++ setup (also these compiler, git, cmake, `cpk` are in `%PATH%`)
2. Having `cmake/modules/.gitkeep` - `cpk` assumes `cmake/cpk_modules` exists (or warn) and expect to overwrite the FindXXX.cmake if applicable
3. `cpk` assumes there are version constraints (exact or ranged, **`master` should not be supported**)

### Supporting libraries
- [ ] [SFML](https://github.com/dirkarnez/sfml-prebuilt), WIP
      - `v2.5.1`
- [ ] [cycfi/elements](https://github.com/dirkarnez/elements-prebuilt), WIP
    - `cd7106afe85fe99c28a3d7314a7ea5d3bd50e9c84`
- [ ] [dacap/clip](https://github.com/dirkarnez/clip-prebuilt), WIP
    - `v1.4`
- [ ] [GiovanniDicanio/WinReg](https://github.com/dirkarnez/WinReg-prebuilt), WIP
    - `v4.1.2`
- [ ] [sciplot/sciplot](https://github.com/dirkarnez/sciplot-prebuilt), WIP
    - `v0.2.2`
- [ ] Boost
- [ ] [LibreDWG/libredwg](https://github.com/dirkarnez/libredwg-prebuilt), WIP
    - `ce32962da3f247937f6355f80a113ea716016e998`
- [ ] [JagPDF - Downloads](http://www.jagpdf.org/downloads.htm)
    - [MAGANER/PdfBuilder: markdown-like language to generate pdf documents.](https://github.com/MAGANER/PdfBuilder)
- [ ] Qt
- [ ] [verilator/verilator: Verilator open-source SystemVerilog simulator and lint system](https://github.com/verilator/verilator)
- [ ] [llvm-project/pstl at main · llvm/llvm-project](https://github.com/llvm/llvm-project/tree/main/pstl)
- [ ] [armadillo](https://github.com/dirkarnez/armadillo-prebuilt), WIP
    - `v11.1.1`
- [ ] [**alecthomas/entityx**](https://github.com/alecthomas/entityx)
- [ ] [jackaudio/jack2: jack2 codebase](https://github.com/jackaudio/jack2)
- [ ] [PortAudio/portaudio: PortAudio is a cross-platform, open-source C language library for real-time audio input and output.](https://github.com/PortAudio/portaudio)
- [ ] [alsa-project/alsa-lib: The Advanced Linux Sound Architecture (ALSA) - library](https://github.com/alsa-project/alsa-lib)
- [ ] [catchorg/Catch2](https://github.com/dirkarnez/catch2-prebuilt), WIP
    - `v3.1.0`
- [ ] [fltk/fltk](https://github.com/dirkarnez/fltk-prebuilt), WIP
    - `v1.3.8`
- [ ] [craigsapp/midifile](https://github.com/dirkarnez/midifile-prebuilt), WIP
    - `cf7ded879139061c41631797e5b735d6d5ceee23e`
- [ ] [FreeGLUTProject/freeglut](https://github.com/dirkarnez/freeglut-prebuilt), WIP
    - `v3.2.2`
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
- [ ] glfw
- [ ] [OGRECave/ogre-next](https://github.com/OGRECave/ogre-next)
- [ ] QuantLib
- [ ] UnQLite
- [ ] sqlite
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
    - `v4.5.5`
- [ ] Magick++
- [ ] https://github.com/facebook/folly
- [ ] Simd
- [ ] gettext
- [ ] cpp-netlib
- [ ] [pybind/pybind11: Seamless operability between C++11 and Python](https://github.com/pybind/pybind11)
- [ ] https://github.com/snakster/cpp.react
- [ ] IBM ICU
- [ ] libiconv
- [ ] DataFrame
- [ ] PortAudio
- [ ] [FFTW/fftw3](https://github.com/dirkarnez/fftw3-prebuilt), WIP
    - `v3.3.10`
- [ ] [oatpp/oatpp](https://github.com/dirkarnez/oatpp-prebuilt), WIP
    - `v1.3.0`
- [ ] [madler/zlib](https://github.com/dirkarnez/zlib-prebuilt), WIP
    - `v1.2.12`
- [ ] OpenMP
- [ ] https://github.com/steven-varga/h5cpp
- [ ] gmplib
- [ ] mpfr
- [ ] GStreamer
- [ ] [libsdl-org/SDL](https://github.com/dirkarnez/sdl-prebuilt), WIP
    - `v2.0.2`
- [ ] TagLib
- [ ] cpprestsdk
- [ ] https://github.com/google/brotli
- [ ] Google tests
- [ ] gRPC
- [ ] [davisking/dlib](https://github.com/dirkarnez/dlib-prebuilt), WIP
    - `v19.24`
- [ ] darknet
- [ ] https://github.com/GHamrouni/Recommender
- [ ] [libeigen / eigen · GitLab](https://github.com/dirkarnez/eigen-prebuilt), WIP
    - `v3.4.0`
- [ ] https://github.com/softwareQinc/qpp
- [ ] https://github.com/libevent/libevent
- [ ] https://github.com/the-tcpdump-group/libpcap
- [ ] https://github.com/protocolbuffers/protobuf
- [ ] https://github.com/USCiLab/cereal
- [ ] ffmpeg
- [ ] https://github.com/jakogut/tinyvm
- [ ] CGAL
- [ ] CUDA
- [ ] [jbeder/yaml-cpp](https://github.com/dirkarnez/yaml-cpp-prebuilt), WIP
    - `v0.7.0`
- [ ] lua
- [ ] Bcrypt
- [ ] curl
- [ ] OpenSSL
    - https://github.com/janbar/openssl-cmake
- [ ] RaftLib
- [ ] GNU Scientific Library
- [ ] fontconfig
- [ ] [electronicarts/EASTL: EASTL stands for Electronic Arts Standard Template Library. It is an extensive and robust implementation that has an emphasis on high performance.](https://github.com/electronicarts/EASTL)
- [ ] [ETLCPP/etl: Embedded Template Library](https://github.com/ETLCPP/etl)
- [ ] [jfalcou/eve: Expressive Vector Engine - SIMD in C++ Goes Brrrr](https://github.com/jfalcou/eve)
- [ ] [g-truc/glm: OpenGL Mathematics (GLM)](https://github.com/g-truc/glm)
- [ ] [HDFGroup/hdf5: Official HDF5® Library Repository](https://github.com/HDFGroup/hdf5)  
- [ ] [oneapi-src/oneTBB: oneAPI Threading Building Blocks (oneTBB)](https://github.com/oneapi-src/oneTBB)
- [ ] [modm-io/avr-libstdcpp: Subset of the C++ standard library for AVR targets](https://github.com/modm-io/avr-libstdcpp)
- [ ] [webview/webview: Tiny cross-platform webview library for C/C++/Golang. Uses WebKit (Gtk/Cocoa) and Edge (Windows)](https://github.com/webview/webview)
- [ ] [andrew-gresyk/HFSM2: High-Performance Hierarchical Finite State Machine Framework](https://github.com/andrew-gresyk/HFSM2)
- [ ] [agenium-scale/nsimd: Agenium Scale vectorization library for CPUs and GPUs](https://github.com/agenium-scale/nsimd/)
- [ ] [jfalcou/eve: Expressive Vector Engine - SIMD in C++ Goes Brrrr](https://github.com/jfalcou/eve)
- [ ] [libuv/libuv: Cross-platform asynchronous I/O](https://github.com/libuv/libuv)
- [ ] [dmlc/xgboost: Scalable, Portable and Distributed Gradient Boosting (GBDT, GBRT or GBM) Library, for Python, R, Java, Scala, C++ and more. Runs on single machine, Hadoop, Spark, Dask, Flink and DataFlow](https://github.com/dmlc/xgboost)
- [ ] cairo
- [ ] [hselasky/libumidi: MIDI file and interface library](https://github.com/hselasky/libumidi)
- [ ] [hselasky/midipp: MIDI Player Pro](https://github.com/hselasky/midipp)
- [ ] [hselasky/jack_umidi: Raw MIDI to JACK daemon](https://github.com/hselasky/jack_umidi)
- [ ] [harfbuzz/harfbuzz: HarfBuzz text shaping engine](https://github.com/harfbuzz/harfbuzz)
- [ ] [freetype/freetype: Official mirror of https://gitlab.freedesktop.org/freetype/freetype](https://github.com/freetype/freetype)
