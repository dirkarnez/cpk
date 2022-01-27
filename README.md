cpk
===
A real stupid C/C++ package manager for own studying (intended not to use other C++ package managers available).  
**And YES! `cpk` sounds exactly like swearing in Cantonese-mixing-English. hahaha :P**

### Convention / Rules / Targets
1. `cpk` must be portable (no registry, no system directory access)
    - store the fetched C++ packages in its directory
2. `cpk` must support different compilers
    - For Windows, mingw and MSVC are a must. 
3. `TARGET_COMPILER` should be well-defined and consistent
4. `TARGET_TAG` should stick to the original versioning practice (eg. some may version "v2.5.1", some may version "2.5.1")
5. `TARGET_LIBRARY` is the library name, it should follow the original name
6. There should be a centralized repository store the information of prebuilt libraries (exactly like `npm`)
7. The creation of prebuilt libraries and its repository should be automated in GitHub action

### Supporting libraries
- [ ] [SFML](https://github.com/dirkarnez/sfml-prebuilt), WIP
  - v2.5.1
- [ ] Boost
- [ ] Qt
- [ ] fltk
- [ ] glfw
- [ ] QuantLib
- [ ] unqlite
- [ ] DataFrame
- [ ] portaudio
- [ ] oatpp
- [ ] curl
