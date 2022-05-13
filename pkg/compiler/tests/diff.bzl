def _diff_test_impl(ctx):
    sh = ctx.actions.declare_file(ctx.attr.name)
    ctx.actions.write(
        sh,
        "#!/bin/sh -e\n%s -s -extra %s > after\ndiff -u %s after\n" % (
            ctx.executable._gofumpt.short_path,
            ctx.file.after.short_path,
            ctx.file.before.short_path,
        ),
    )
    return [DefaultInfo(
        executable = sh,
        runfiles = ctx.runfiles(files = [
            ctx.executable._gofumpt,
            ctx.file.before,
            ctx.file.after,
        ]),
    )]

diff_test = rule(
    implementation = _diff_test_impl,
    attrs = {
        "_gofumpt": attr.label(
            default = "@cc_mvdan_gofumpt//:gofumpt",
            executable = True,
            cfg = "host",
        ),
        "before": attr.label(allow_single_file = True),
        "after": attr.label(allow_single_file = True),
    },
    test = True,
)
