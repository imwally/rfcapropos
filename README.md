rfcapropos
==========
Like apropos(1) but for RFCs.

```
rfcapropos(1)                                                               rfcapropos(1)

NAME
        rfcapropos - search for RFCs by keyword or phrase

SYNOPSIS
        rfcapropos searches for Request For Comments by keywords or phrases and
        displays the results on the standard output. 

NOTES
        An internet connection is required. Search terms are sent to the
        rfcsearch API at https://rfcsearch-gorun.rhcloud.com which in turn are
        sent to RFC Editor's search page at
        https://www.rfc-editor.org/search/rfc_search.php. Only alphanumerical
        characters are allowed. For example, instead of searching for "SSL 3.0"
        try the phrase "SSL 3" or "tls 1" instead of "tls1.2".

AUTHOR
        Wally Jones <wally@imwally.net>

SEE ALSO
        apropos(1)

                                    September 13, 2016                      rfcapropos(1)
```

Install
-------
`$ go get github.com/imwally/rfcapropos`
