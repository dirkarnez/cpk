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
8. The creation of prebuilt libraries and their repository should be automated in GitHub action and repository template
9. `cpk` and the supporting library repositories should also support listing include paths for easier development in VSCode
10. For C/C++ projects using `cpk`, there should not be any configuration files needed - `cpk` should only depend on (read / write) explicit `cmake` file complying project onfiguration (`find_package` in `CMakeLists.txt`)
11. `cpk` should leave room for modifying third-party libraries (for example: editing original `CMakeLists.txt` and may / may not submit PR in a **fork repo** and rebuilding prebuilts, including ad-hoc /private redirection of libraries
13. `cpk` should resolve dependency-related issues like nested dependency
14. `cpk` should give hints (using CMake macros?) on further setup for the projects using the installed libraries
15. `cpk` should be unobtrusive, play well with unsupported / local libraries - `cpk` should have a dedicated folder for its CMake modules fetching and configuration

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
- [ ] Boost
- [ ] Qt
- [ ] fltk
- [ ] CImg
- [ ] glfw
- [ ] QuantLib
- [ ] UnQLite
- [ ] sqlite
- [ ] libjpeg
- [ ] libpng
- [ ] OpenCV
- [ ] Magick++
- [ ] Simd
- [ ] gettext
- [ ] cpp-netlib
- [ ] [pybind/pybind11: Seamless operability between C++11 and Python](https://github.com/pybind/pybind11)
- [ ] https://github.com/snakster/cpp.react
- [ ] IBM ICU
- [ ] libiconv
- [ ] DataFrame
- [ ] PortAudio
- [ ] oatpp
- [ ] OpenMP
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
