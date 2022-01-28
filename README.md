cpk
===
A real stupid C/C++ package manager for own studying (intended not to use other C++ package managers available).  
**And YES! `cpk` sounds exactly like swearing in Cantonese-mixing-English. hahaha :P**

### Convention / Rules / Targets
1. `cpk` must be portable (no registry, no system directory access, no `%PATH%` modifying, no `pkg-config`, no crazy path-searching)
    - store the fetched C++ packages in its directory instead of %USERPROFILE% or `~`
2. `cpk` must support different compilers
    - For Windows, mingw and MSVC are a must. 
3. `TARGET_COMPILER` should be well-defined and consistent, with compiler version
4. `TARGET_TAG` should stick to the original versioning practice (eg. some may version "v2.5.1", some may version "2.5.1")
5. `TARGET_LIBRARY` is the library name, it should follow the original name
6. There should be a centralized repository store the information of prebuilt libraries (exactly like `npm`)
7. The creation of prebuilt libraries and their repository should be automated in GitHub action
8. `cpk` and the supporting library repositories should also support listing include paths for easier development in VSCode
9. For C/C++ projects using `cpk`, there should not be any configuration files needed - `cpk` should only depend on (read / write) explicit `cmake` file complying project configuration (`find_package in CMakeLists.txt`)
10. `cpk` should leave room for modifying third-party libraries and rebuilding prebuilts
11. `cpk` should resolve dependency-related issues like nested dependency
12. `cpk` should give hints on further setup for the projects uisng the installed libraries

### Supporting libraries
- [ ] [SFML](https://github.com/dirkarnez/sfml-prebuilt), WIP
  - v2.5.1
- [ ] Boost
- [ ] Qt
- [ ] fltk
- [ ] glfw
- [ ] QuantLib
- [ ] UnQLite
- [ ] sqlite
- [ ] libjpeg
- [ ] OpenCV
- [ ] Magick++
- [ ] Simd
- [ ] gettext
- [ ] cpp-netlib
- [ ] https://github.com/snakster/cpp.react
- [ ] IBM ICU
- [ ] libiconv
- [ ] DataFrame
- [ ] PortAudio
- [ ] oatpp
- [ ] OpenMP
- [ ] gmplib
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
- [ ] yaml-cpp
- [ ] lua
- [ ] Bcrypt
- [ ] curl
- [ ] OpenSSL
- [ ] RaftLib
- [ ] GNU Scientific Library
- [ ] elements
