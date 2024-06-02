//go:build archivematica

package config

func init() {
	siegfried.signature = "archivematica.sig"
	identifier.name = "archivematica"
	identifier.extend = []string{"archivematica-fmt2.xml", "archivematica-fmt3.xml", "archivematica-fmt4.xml", "archivematica-fmt5.xml"}
}

func defaultHome() string {
	return "/usr/share/siegfried"
}
