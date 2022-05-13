XDRLibrary = provider()

def _go_xdr_library_impl(ctx):
    all_srcs = depset(
        direct = [ctx.file.src],
        transitive = [dep[XDRLibrary].srcs for dep in ctx.attr.deps],
    )
    out = ctx.actions.declare_file(ctx.attr.name + ".go")
    ctx.actions.run(
        inputs = all_srcs,
        outputs = [out],
        executable = ctx.executable._xdr_compiler,
        arguments = [ctx.file.src.path, out.path],
    )
    return [
        DefaultInfo(files = depset([out])),
        XDRLibrary(srcs = all_srcs),
    ]

go_xdr_library = rule(
    implementation = _go_xdr_library_impl,
    attrs = {
        "_xdr_compiler": attr.label(
            default = "@com_github_buildbarn_go_xdr//cmd/xdr_compiler",
            executable = True,
            cfg = "host",
        ),
        "src": attr.label(allow_single_file = True),
        "deps": attr.label_list(providers = [XDRLibrary]),
    },
)
