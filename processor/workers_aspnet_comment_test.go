// SPDX-License-Identifier: MIT

package processor

import (
	"testing"
)

// The ASP.NET server side comment <%-- --%> was paired with the HTML closer
// -->, so the opener never closed and every line after it was counted as a
// comment. The <!-- --> pair in the same entry is unaffected.

func TestCountStatsAspNetServerSideComment(t *testing.T) {
	ProcessConstants()
	fileJob := FileJob{
		Language: "ASP.NET",
	}

	fileJob.SetContent(`<%@ Page Language="C#" %>
<%-- not sent to the browser --%>
<html>
<body>
<form id="form1" runat="server" />
</body>
</html>`)

	CountStats(&fileJob)

	if fileJob.Lines != 7 {
		t.Errorf("Expected 7 lines got %d", fileJob.Lines)
	}
	if fileJob.Code != 6 {
		t.Errorf("Expected 6 code got %d", fileJob.Code)
	}
	if fileJob.Comment != 1 {
		t.Errorf("Expected 1 comment got %d", fileJob.Comment)
	}
	if fileJob.Blank != 0 {
		t.Errorf("Expected 0 blanks got %d", fileJob.Blank)
	}
}

func TestCountStatsAspNetMultiLineServerSideComment(t *testing.T) {
	ProcessConstants()
	fileJob := FileJob{
		Language: "ASP.NET",
	}

	fileJob.SetContent(`<%--
  old markup, left here on purpose
--%>
<html>
<!-- browser comment -->
</html>`)

	CountStats(&fileJob)

	if fileJob.Lines != 6 {
		t.Errorf("Expected 6 lines got %d", fileJob.Lines)
	}
	if fileJob.Code != 2 {
		t.Errorf("Expected 2 code got %d", fileJob.Code)
	}
	if fileJob.Comment != 4 {
		t.Errorf("Expected 4 comments got %d", fileJob.Comment)
	}
	if fileJob.Blank != 0 {
		t.Errorf("Expected 0 blanks got %d", fileJob.Blank)
	}
}
