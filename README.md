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
- [ ] Qt
- [ ] [armadillo](https://github.com/dirkarnez/armadillo-prebuilt), WIP
    - `v11.1.1`
- [ ] [fltk/fltk](https://github.com/dirkarnez/fltk-prebuilt), WIP
    - `v1.3.8`
- [ ] CImg
- [ ] [antirez/sds](https://github.com/dirkarnez/sds-prebuilt), **non-cmake** repo, **WIP**
- [ ] https://github.com/mlpack/mlpack
- [ ] glfw
- [ ] QuantLib
- [ ] UnQLite
- [ ] sqlite
- [ ] [laserpants/dotenv-cpp](https://github.com/dirkarnez/dotenv-cpp-prebuilt), WIP
    - `cca1d543f78702921dbb398c865806b71dedfd2eb`
- [ ] [Armadillo: C++ library for linear algebra & scientific computing](http://arma.sourceforge.net/)
- [ ] libjpeg
- [ ] libpng
- [ ] [qgis/QGIS: QGIS is a free, open source, cross platform (lin/win/mac) geographical information system (GIS)](https://github.com/qgis/QGIS)
- [ ] [graphql/libgraphqlparser: A GraphQL query parser in C++ with C and C++ APIs](https://github.com/graphql/libgraphqlparser)
- [ ] OpenCV
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
- [ ] [oatpp/oatpp](https://github.com/dirkarnez/oatpp-prebuilt), WIP
    - `v1.3.0`
- [ ] OpenMP
- [ ] https://github.com/steven-varga/h5cpp
- [ ] gmplib
- [ ] mpfr
- [ ] GStreamer
- [ ] SDL
- [ ] TagLib
- [ ] cpprestsdk
- [ ] https://github.com/google/brotli
- [ ] Google tests
- [ ] gRPC
- [ ] Dlib 
- [ ] darknet
- [ ] https://github.com/GHamrouni/Recommender
- [ ] Eigen
- [ ] https://github.com/softwareQinc/qpp
- [ ] https://github.com/libevent/libevent
- [ ] https://github.com/the-tcpdump-group/libpcap
- [ ] FFTW
- [ ] https://github.com/protocolbuffers/protobuf
- [ ] https://github.com/USCiLab/cereal
- [ ] ffmpeg
- [ ] https://github.com/jakogut/tinyvm
- [ ] CGAL
- [ ] ZLib
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
- [ ] cairo
- [ ] [hselasky/libumidi: MIDI file and interface library](https://github.com/hselasky/libumidi)
- [ ] [hselasky/midipp: MIDI Player Pro](https://github.com/hselasky/midipp)
- [ ] [hselasky/jack_umidi: Raw MIDI to JACK daemon](https://github.com/hselasky/jack_umidi)
- [ ] [libjpeg-turbo/libjpeg-turbo: Main libjpeg-turbo repository](https://github.com/libjpeg-turbo/libjpeg-turbo)
- [ ] [harfbuzz/harfbuzz: HarfBuzz text shaping engine](https://github.com/harfbuzz/harfbuzz)
- [ ] [freetype/freetype: Official mirror of https://gitlab.freedesktop.org/freetype/freetype](https://github.com/freetype/freetype)
