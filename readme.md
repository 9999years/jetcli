# Jet CLI

A command-line interface for the [Jet template engine][jet].

Source code resides in the `jet` directory so that the binary name is `jet`.

    Usage: jet [options] [-template] TEMPLATE_NAME
      -dir string
            The directory to search for templates in (default "./")
      -template string
            The filename of the template to render
    NOTE: TEMPLATE_NAME can be given either as a named flag or positionally as
    the last argument

[jet]: https://github.com/CloudyKit/jet
