# Build tensorflow on local

## Mac OS

**config**

```shell
$ cd tensorflow
$ ./configure
```

**clean**

```shell
$ bazel clean
```

**build pip package (CPU)**

```shell
$ bazel build --config=opt //tensorflow/tools/pip_package:build_pip_package
$ bazel-bin/tensorflow/tools/pip_package/build_pip_package /tmp/tensorflow_pkg
```

Add `--config=cuda` with GPU support.

And with `verbose_explanations` for more details:

```shell
$ bazel build --explain log.txt --verbose_explanations=true -s --config=opt //tensorflow/tools/pip_package:build_pip_package
```

**install pip package**

```shell
$ sudo pip install --upgrade --no-deps --force-reinstall /tmp/tensorflow_pkg/tensorflow-1.0.0-cp27-cp27m-macosx_10_11_x86_64.whl
```

- Note: replace `1.0.0 *** ` by your tensorflow version

[force pip to reinstall the current version](https://stackoverflow.com/a/27254355/3167471)

## Common installation problems

**ImportError: No module named pywrap_tensorflow_internal**

cd out of the `tensorflow` directory

[ref](https://stackoverflow.com/a/35963479/3167471)
