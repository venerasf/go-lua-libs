function TestRequireModule(t)
    modules = {
        "argparse",
        "base64",
        "cert_util",
        "chef",
        "cmd",
        "crypto",
        "db",
        "filepath",
        "goos",
        "humanize",
        "inspect",
        "ioutil",
        "json",
        "log",
        "pb",
        "plugin",
        "pprof",
        "regexp",
        "runtime",
        "shellescape",
        "stats",
        "storage",
        "strings",
        "tac",
        "tcp",
        "template",
        "time",
        "xmlpath",
        "yaml",
    }
    for _, module in ipairs(modules) do
        t:Run(module, function(t)
            require(module)
        end)
    end
end

