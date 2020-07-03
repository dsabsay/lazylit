package main

var Css = `
/*--------------------- Layout and Typography ----------------------------*/
body {
  font-family: 'Palatino Linotype', 'Book Antiqua', Palatino, FreeSerif, serif;
  font-size: 15px;
  line-height: 22px;
  color: #252519;
  margin: 0; padding: 0;
}
hr {
    border: 0;
    border-top: 1px solid rgba(0, 0, 0, 0.2);
    margin-bottom: 1rem;
}
footer {
    color: rgba(0, 0, 0, 0.5);
}
a {
  color: #261a3b;
}
  a:visited {
    color: #261a3b;
  }
p {
  margin: 0 0 15px 0;
}
p.footnote {
    font-size: 0.6rem;
}
h1, h2, h3, h4, h5, h6 {
  margin: 0px 0 15px 0;
}
  h1 {
    margin-top: 40px;
  }
#container {
  position: relative;
}
#background {
  position: fixed;
  top: 0; left: 525px; right: 0; bottom: 0;
  background: #f5f5ff;
  border-left: 1px solid #e5e5ee;
  z-index: -1;
}
#content {
    padding: 10px 25px 1px 50px;
    width: 465px;
}
#jump_to, #jump_page {
  background: white;
  -webkit-box-shadow: 0 0 25px #777; -moz-box-shadow: 0 0 25px #777;
  -webkit-border-bottom-left-radius: 5px; -moz-border-radius-bottomleft: 5px;
  font: 10px Arial;
  text-transform: uppercase;
  cursor: pointer;
  text-align: right;
}
#jump_to, #jump_wrapper {
  position: fixed;
  right: 0; top: 0;
  padding: 5px 10px;
}
  #jump_wrapper {
    padding: 0;
    display: none;
  }
    #jump_to:hover #jump_wrapper {
      display: block;
    }
    #jump_page {
      padding: 5px 0 3px;
      margin: 0 0 25px 25px;
    }
      #jump_page .source {
        display: block;
        padding: 5px 10px;
        text-decoration: none;
        border-top: 1px solid #eee;
      }
        #jump_page .source:hover {
          background: #f5f5ff;
        }
        #jump_page .source:first-child {
        }
th {
    font-weight: normal;
}
table td {
  border: 0;
  outline: 0;
}
  td.docs, th.docs {
    max-width: 450px;
    min-width: 450px;
    min-height: 5px;
    padding: 10px 25px 1px 50px;
    overflow-x: hidden;
    vertical-align: top;
    text-align: left;
  }
    .docs pre {
      margin: 15px 0 15px;
      padding-left: 15px;
    }
    .docs p tt, .docs p code {
      background: #f8f8ff;
      border: 1px solid #dedede;
      font-size: 12px;
      padding: 0 0.2em;
    }
    .pilwrap {
      position: relative;
    }
      .pilcrow {
        font: 12px Arial;
        text-decoration: none;
        color: #454545;
        position: absolute;
        top: 3px; left: -20px;
        padding: 1px 2px;
        opacity: 0;
        -webkit-transition: opacity 0.2s linear;
      }
        td.docs:hover .pilcrow {
          opacity: 1;
        }
  td.code, th.code {
    padding: 14px 15px 16px 25px;
    width: 100%;
    vertical-align: top;
    background: #f5f5ff;
    border-left: 1px solid #e5e5ee;
  }
    pre, tt, code {
      font-size: 12px; line-height: 18px;
      font-family: Menlo, Monaco, Consolas, "Lucida Console", monospace;
      margin: 0; padding: 0;
    }


/*---------------------- Syntax Highlighting -----------------------------*/
td.linenos { background-color: #f0f0f0; padding-right: 10px; }
span.lineno { background-color: #f0f0f0; padding: 0 5px 0 5px; }
body .hll { background-color: #ffffcc }
body .c { color: #408080; font-style: italic }  /* Comment */
body .err { border: 1px solid #FF0000 }         /* Error */
body .k { color: #954121 }                      /* Keyword */
body .o { color: #666666 }                      /* Operator */
body .cm { color: #408080; font-style: italic } /* Comment.Multiline */
body .cp { color: #BC7A00 }                     /* Comment.Preproc */
body .c1 { color: #408080; font-style: italic } /* Comment.Single */
body .cs { color: #408080; font-style: italic } /* Comment.Special */
body .gd { color: #A00000 }                     /* Generic.Deleted */
body .ge { font-style: italic }                 /* Generic.Emph */
body .gr { color: #FF0000 }                     /* Generic.Error */
body .gh { color: #000080; font-weight: bold }  /* Generic.Heading */
body .gi { color: #00A000 }                     /* Generic.Inserted */
body .go { color: #808080 }                     /* Generic.Output */
body .gp { color: #000080; font-weight: bold }  /* Generic.Prompt */
body .gs { font-weight: bold }                  /* Generic.Strong */
body .gu { color: #800080; font-weight: bold }  /* Generic.Subheading */
body .gt { color: #0040D0 }                     /* Generic.Traceback */
body .kc { color: #954121 }                     /* Keyword.Constant */
body .kd { color: #954121; font-weight: bold }  /* Keyword.Declaration */
body .kn { color: #954121; font-weight: bold }  /* Keyword.Namespace */
body .kp { color: #954121 }                     /* Keyword.Pseudo */
body .kr { color: #954121; font-weight: bold }  /* Keyword.Reserved */
body .kt { color: #B00040 }                     /* Keyword.Type */
body .m { color: #666666 }                      /* Literal.Number */
body .s { color: #219161 }                      /* Literal.String */
body .na { color: #7D9029 }                     /* Name.Attribute */
body .nb { color: #954121 }                     /* Name.Builtin */
body .nc { color: #0000FF; font-weight: bold }  /* Name.Class */
body .no { color: #880000 }                     /* Name.Constant */
body .nd { color: #AA22FF }                     /* Name.Decorator */
body .ni { color: #999999; font-weight: bold }  /* Name.Entity */
body .ne { color: #D2413A; font-weight: bold }  /* Name.Exception */
body .nf { color: #0000FF }                     /* Name.Function */
body .nl { color: #A0A000 }                     /* Name.Label */
body .nn { color: #0000FF; font-weight: bold }  /* Name.Namespace */
body .nt { color: #954121; font-weight: bold }  /* Name.Tag */
body .nv { color: #19469D }                     /* Name.Variable */
body .ow { color: #AA22FF; font-weight: bold }  /* Operator.Word */
body .w { color: #bbbbbb }                      /* Text.Whitespace */
body .mf { color: #666666 }                     /* Literal.Number.Float */
body .mh { color: #666666 }                     /* Literal.Number.Hex */
body .mi { color: #666666 }                     /* Literal.Number.Integer */
body .mo { color: #666666 }                     /* Literal.Number.Oct */
body .sb { color: #219161 }                     /* Literal.String.Backtick */
body .sc { color: #219161 }                     /* Literal.String.Char */
body .sd { color: #219161; font-style: italic } /* Literal.String.Doc */
body .s2 { color: #219161 }                     /* Literal.String.Double */
body .se { color: #BB6622; font-weight: bold }  /* Literal.String.Escape */
body .sh { color: #219161 }                     /* Literal.String.Heredoc */
body .si { color: #BB6688; font-weight: bold }  /* Literal.String.Interpol */
body .sx { color: #954121 }                     /* Literal.String.Other */
body .sr { color: #BB6688 }                     /* Literal.String.Regex */
body .s1 { color: #219161 }                     /* Literal.String.Single */
body .ss { color: #19469D }                     /* Literal.String.Symbol */
body .bp { color: #954121 }                     /* Name.Builtin.Pseudo */
body .vc { color: #19469D }                     /* Name.Variable.Class */
body .vg { color: #19469D }                     /* Name.Variable.Global */
body .vi { color: #19469D }                     /* Name.Variable.Instance */
body .il { color: #666666 }                     /* Literal.Number.Integer.Long */
`

var ABOUT_HTML = `
<!DOCTYPE html>

<html>
<head>
    <title>About lazylit</title>
  <meta http-equiv="content-type" content="text/html; charset=UTF-8">
  <link rel="stylesheet" media="all" href="gocco.css" />
</head>

<body>
  <div id="container">
    <div id="background"></div>
    <div id="content">
        <h1> What is lazylit? </h1>
        <p>
            Lazylit is a collection of heavily documented source code files. Each page of documentation hosted here is written for a specific revision/commit of a source code file that lives somewhere else (i.e. a repository in GitHub).
        </p>
        <p>
            Follow the links<sup>*</sup> below to view available documentation:
        </p>
        <ul>
            {{ range . }}
            <li>
                <a href="{{ . }}/index.html">
                {{ . }}
                </a>
            </li>
            {{ end }}
        </ul>
        <p class="footnote">
            <sup>*</sup> They link to "redirection pages" which provides a consistent identifier and landing page even if the source code file changes names over time.
        </p>
        <hr>
        <h2>Motivation</h2>
        <p>
            I wanted a simple way to share extensive code comments without the ongoing maintenance burden and inevitable code/comment drift that is characteristic of traditional code commenting practice. 
        </p>
        <p>
            I've often heard advice that usually sounds like "don't use advanced features of a language/tool, because it makes code harder to understand".
            With a tool like GNU <code>make</code>, a lot of these "advanced" features can replace dozens of lines of bespoke shell scripts.
            Following this advice then turns an opportunity to learn a common tool <i>well</i> into a grueling slog through a coworker's (or your own) buggy mess.
            Instead of avoiding certain features and tools, I believe that a well-written set of notes (with links to relevant documentation) can go a long way in improving code understandability.
            There are thousands of programming tools and languages; most people only know a few.
            For <code>make</code> in particular, my (limited) experience has shown me that programmers without experience in C are likely to be unfamiliar with <code>make</code> in general.
            Expecting that all code should be immediately understandable to anyone at first glance is unrealistic.
            But we can certainly leave clues behind us that point readers <i>toward</i> understanding.
        </p>
        <p>
            But fancy tool/language features are not the only thing worth documenting in this way.
            <i>Project</i>- and <i>team</i>-specific conventions and patterns are worth documenting as well.
            For example, if all your team's projects have a Makefile with a similar structure, an explanation of that structure and pattern could accelerate a new developer's ability to read and understand the team's projects.
        </p>
        <p>
            Approaching existing code bases is one of the hardest things I've had to do as a programmer. This project as an attempt to address that challenge. I hope it will make the lives of my coworkers easier by making it easier for new developers to learn about the conventions and tricks used in codebases I work on.
        </p>
        <h2> Why not store the explanation in the repo, next to the source itself? </h2>
        <p>
            Storing an explanation next to the code (either in comments, as separate files, or using a proper <a href="https://en.wikipedia.org/wiki/Literate_programming">literate programming</a> tool) comes with the expectation that it is always kept up-to-date with the source.
            This incurs a maintenance burden, which I felt would be too cumbersome and in fact, unnecessary.
            The focus should be on explaining the <i>patterns</i> and the <i>features</i> or <i>paradigms</i> being used; the things that will help a developer read the code itself.
            Thus, an explanation written against a version of the code slightly out of date should still be useful, because they can apply what they learn to the newer code.
        </p>
        <h2>Acknowledgements</h2>
        <p>
            I borrowed heavily from the <a href="http://ashkenas.com/docco/">Docco</a> project and its derivatives. I am especially grateful for Nikhil Marathe's golang port of Docco, <a href="https://github.com/nikhilm/gocco">gocco</a>, from which I borrowed (i.e. copied) a large portion of this code.
        </p>
        <hr>
        <footer>
            <p>Lazylit was created by Daniel Sabsay and is under the MIT License.</p>
        </footer>
    </div>
  </div>
</body>
</html>
`

var INDEX_HTML = `
<!DOCTYPE html>

<html>
<head>
    <title>{{ .ArtifactName }}</title>
  <meta http-equiv="content-type" content="text/html; charset=UTF-8">
  <link rel="stylesheet" media="all" href="../gocco.css" />
</head>

<body>
  <div id="container">
    <div id="background"></div>
    <div id="content">
        <h1> {{ .ArtifactName }} </h1>
        <p> Notes are available for these revisions (commits): </p>
        <ul>
            {{ range .Snapshots }}
            <li>
                <a href="{{ .Destination | base }}">
                {{ .CommitDateString }} ({{ .SourceFileName }})
                </a>
            </li>
            {{ end }}
        </ul>
    </div>
  </div>
</body>
</html>
`

var HTML = `
<!DOCTYPE html>

<html>
<head>
    <title>{{ .Title }}</title>
  <meta http-equiv="content-type" content="text/html; charset=UTF-8">
  <link rel="stylesheet" media="all" href="../gocco.css" />
</head>
<body>
  <div id="container">
    <div id="background"></div>
    {{ if .Multiple }}
      <div id="jump_to">
        Other revisions &hellip;
        <div id="jump_wrapper">
          <div id="jump_page">
              {{ range .OtherRevisions }}
              <a class="source" href="{{ .Destination | base }}">
                  {{ .CommitDateString }}
              </a>
              {{ end }}
          </div>
        </div>
      </div>
    {{ end }}
    <table cellpadding="0" cellspacing="0">
      <thead>
        <tr>
          <th class="docs">
            <h1>
                {{ .Title }}
            </h1>
            <p> <i>
                Viewing notes written by {{ .Snapshot.DocAuthor }} for {{ .Snapshot.SourceFileName }} at revision <a href="{{ .Snapshot.SourceLink }}">{{ .Snapshot.Commit }} ({{ .Snapshot.CommitDateString }})</a>. Select other revisions via the menu to the right.
            </i> </p>
          </th>
          <th class="code">
          </th>
        </tr>
      </thead>
      <tbody>
          {{ range .Sections }}
          <tr id="section-{{ .Index }}">
            <td class="docs">
              <div class="pilwrap">
                  <a class="pilcrow" href="#section-{{ .Index }}">&#182;</a>
              </div>
                {{ .DocsHTML }}
            </td>
            <td class="code">
                {{ .CodeHTML }}
            </td>
          </tr>
          {{ end }}
      </tbody>
    </table>
  </div>
</body>
</html>
`
