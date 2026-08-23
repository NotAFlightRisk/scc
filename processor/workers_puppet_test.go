// SPDX-License-Identifier: MIT

package processor

import (
	"testing"
)

// The Puppet language entry carried Ruby's =begin/=end block comment and the
// two byte \" string delimiter. Puppet's block comment is /* */, so a header
// comment was counted as code, and once /* */ is recognised the double quote
// has to work as well or a path glob such as "/etc/nginx/conf.d/*/*.conf"
// opens a comment that runs to the next */ or to EOF.

func TestCountStatsPuppetBlockComment(t *testing.T) {
	ProcessConstants()
	fileJob := FileJob{
		Language: "Puppet",
	}

	fileJob.SetContent(`/*
  Manages the nginx package
*/
class nginx {
  package { 'nginx':
    ensure => installed,
  }
}`)

	CountStats(&fileJob)

	if fileJob.Lines != 8 {
		t.Errorf("Expected 8 lines got %d", fileJob.Lines)
	}
	if fileJob.Code != 5 {
		t.Errorf("Expected 5 code got %d", fileJob.Code)
	}
	if fileJob.Comment != 3 {
		t.Errorf("Expected 3 comments got %d", fileJob.Comment)
	}
	if fileJob.Blank != 0 {
		t.Errorf("Expected 0 blanks got %d", fileJob.Blank)
	}
}

func TestCountStatsPuppetDoubleQuotedGlobNotComment(t *testing.T) {
	ProcessConstants()
	fileJob := FileJob{
		Language: "Puppet",
	}

	fileJob.SetContent(`class nginx::config {
  $sites = "/etc/nginx/conf.d/*/*.conf"
  notify { $sites: }
}`)

	CountStats(&fileJob)

	if fileJob.Lines != 4 {
		t.Errorf("Expected 4 lines got %d", fileJob.Lines)
	}
	if fileJob.Code != 4 {
		t.Errorf("Expected 4 code got %d", fileJob.Code)
	}
	if fileJob.Comment != 0 {
		t.Errorf("Expected 0 comments got %d", fileJob.Comment)
	}
	if fileJob.Blank != 0 {
		t.Errorf("Expected 0 blanks got %d", fileJob.Blank)
	}
}
