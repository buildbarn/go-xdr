XDRLibrary = provider()

def _go_xdr_library_impl(ctx):
    import_root = (
        ctx.file.src.root.path + ctx.label.workspace_root + ctx.attr.strip_import_prefix
    )
    import_root_prefix = import_root
    if import_root_prefix:
        import_root_prefix += "/"
    if not ctx.file.src.path.startswith(import_root_prefix):
        fail(
            "Source path %s does not start with import root %s" %
            (ctx.file.src.path, import_root_prefix),
        )
    short_src_path = ctx.file.src.path[len(import_root_prefix):]

    all_srcs = depset(
        direct = [ctx.file.src],
        transitive = [dep[XDRLibrary].srcs for dep in ctx.attr.deps],
    )
    import_roots = depset(
        direct = [import_root],
        transitive = [dep[XDRLibrary].import_roots for dep in ctx.attr.deps],
    )

    out = ctx.actions.declare_file(ctx.attr.name + ".go")
    args = ctx.actions.args()
    args.add_all(import_roots, before_each = "-I")
    args.add(short_src_path)
    args.add(out.path)
    ctx.actions.run(
        inputs = all_srcs,
        outputs = [out],
        executable = ctx.executable._xdr_compiler,
        arguments = [args],
    )

    return [
        DefaultInfo(files = depset([out])),
        XDRLibrary(srcs = all_srcs, import_roots = import_roots),
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
        "strip_import_prefix": attr.string(default = ""),
    },
)
