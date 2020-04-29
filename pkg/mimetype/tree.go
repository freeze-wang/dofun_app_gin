package mimetype

import "dofun/pkg/mimetype/internal/matchers"

// root is a matcher which passes for any slice of bytes.
// When a matcher passes the check, the children matchers
// are tried in order to find a more accurate mime type.
var root = newNode("application/octet-stream", "", matchers.True,
	sevenZ, zip, pdf, doc, xls, ppt, ps, psd, ogg, png, jpg, gif, webp, exe, elf,
	ar, tar, xar, bz2, fits, tiff, bmp, ico, mp3, flac, midi, ape, musePack, amr,
	wav, aiff, au, mpeg, quickTime, mqv, mp4, webM, threeGP, threeG2, avi, flv,
	mkv, asf, aMp4, m4a, txt, gzip, class, swf, crx, woff, woff2, wasm, shx, dbf,
	dcm,
)

// The list of nodes appended to the root node
var (
	gzip   = newNode("application/gzip", "gz", matchers.Gzip)
	sevenZ = newNode("application/x-7z-compressed", "7z", matchers.SevenZ)
	zip    = newNode("application/zip", "zip", matchers.Zip, xlsx, docx, pptx, epub, jar)
	tar    = newNode("application/x-tar", "tar", matchers.Tar)
	xar    = newNode("application/x-xar", "xar", matchers.Xar)
	bz2    = newNode("application/x-bzip2", "bz2", matchers.Bz2)
	pdf    = newNode("application/pdf", "pdf", matchers.Pdf)
	xlsx   = newNode("application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", "xlsx", matchers.Xlsx)
	docx   = newNode("application/vnd.openxmlformats-officedocument.wordprocessingml.document", "docx", matchers.Docx)
	pptx   = newNode("application/vnd.openxmlformats-officedocument.presentationml.presentation", "pptx", matchers.Pptx)
	epub   = newNode("application/epub+zip", "epub", matchers.Epub)
	jar    = newNode("application/jar", "jar", matchers.Jar)
	doc    = newNode("application/msword", "doc", matchers.Doc)
	ppt    = newNode("application/vnd.ms-powerpoint", "ppt", matchers.Ppt)
	xls    = newNode("application/vnd.ms-excel", "xls", matchers.Xls)
	ps     = newNode("application/postscript", "ps", matchers.Ps)
	psd    = newNode("application/x-photoshop", "psd", matchers.Psd)
	fits   = newNode("application/fits", "fits", matchers.Fits)
	ogg    = newNode("application/ogg", "ogg", matchers.Ogg)
	txt    = newNode("text/plain", "txt", matchers.Txt,
		html, svg, xml, php, js, lua, perl, python, json, rtf, tcl, csv, tsv)
	xml = newNode("text/xml; charset=utf-8", "xml", matchers.Xml,
		x3d, kml, collada, gml, gpx, tcx)
	json      = newNode("application/json", "json", matchers.Json, geoJson)
	csv       = newNode("text/csv", "csv", matchers.Csv)
	tsv       = newNode("text/tab-separated-values", "tsv", matchers.Tsv)
	geoJson   = newNode("application/geo+json", "geojson", matchers.GeoJson)
	html      = newNode("text/html; charset=utf-8", "html", matchers.Html)
	php       = newNode("text/x-php; charset=utf-8", "php", matchers.Php)
	rtf       = newNode("text/rtf", "rtf", matchers.Rtf)
	js        = newNode("application/javascript", "js", matchers.Js)
	lua       = newNode("text/x-lua", "lua", matchers.Lua)
	perl      = newNode("text/x-perl", "pl", matchers.Perl)
	python    = newNode("application/x-python", "py", matchers.Python)
	tcl       = newNode("text/x-tcl", "tcl", matchers.Tcl)
	svg       = newNode("image/svg+xml", "svg", matchers.Svg)
	x3d       = newNode("model/x3d+xml", "x3d", matchers.X3d)
	kml       = newNode("application/vnd.google-earth.kml+xml", "kml", matchers.Kml)
	collada   = newNode("model/vnd.collada+xml", "dae", matchers.Collada)
	gml       = newNode("application/gml+xml", "gml", matchers.Gml)
	gpx       = newNode("application/gpx+xml", "gpx", matchers.Gpx)
	tcx       = newNode("application/vnd.garmin.tcx+xml", "tcx", matchers.Tcx)
	png       = newNode("image/png", "png", matchers.Png)
	jpg       = newNode("image/jpeg", "jpg", matchers.Jpg)
	gif       = newNode("image/gif", "gif", matchers.Gif)
	webp      = newNode("image/webp", "webp", matchers.Webp)
	tiff      = newNode("image/tiff", "tiff", matchers.Tiff)
	bmp       = newNode("image/bmp", "bmp", matchers.Bmp)
	ico       = newNode("image/x-icon", "ico", matchers.Ico)
	mp3       = newNode("audio/mpeg", "mp3", matchers.Mp3)
	flac      = newNode("audio/flac", "flac", matchers.Flac)
	midi      = newNode("audio/midi", "midi", matchers.Midi)
	ape       = newNode("audio/ape", "ape", matchers.Ape)
	musePack  = newNode("audio/musepack", "mpc", matchers.MusePack)
	wav       = newNode("audio/wav", "wav", matchers.Wav)
	aiff      = newNode("audio/aiff", "aiff", matchers.Aiff)
	au        = newNode("audio/basic", "au", matchers.Au)
	amr       = newNode("audio/amr", "amr", matchers.Amr)
	aMp4      = newNode("audio/mp4", "mp4", matchers.AMp4)
	m4a       = newNode("audio/x-m4a", "m4a", matchers.M4a)
	mp4       = newNode("video/mp4", "mp4", matchers.Mp4)
	webM      = newNode("video/webm", "webm", matchers.WebM)
	mpeg      = newNode("video/mpeg", "mpeg", matchers.Mpeg)
	quickTime = newNode("video/quicktime", "mov", matchers.QuickTime)
	mqv       = newNode("video/quicktime", "mqv", matchers.Mqv)
	threeGP   = newNode("video/3gpp", "3gp", matchers.ThreeGP)
	threeG2   = newNode("video/3gpp2", "3g2", matchers.ThreeG2)
	avi       = newNode("video/x-msvideo", "avi", matchers.Avi)
	flv       = newNode("video/x-flv", "flv", matchers.Flv)
	mkv       = newNode("video/x-matroska", "mkv", matchers.Mkv)
	asf       = newNode("video/x-ms-asf", "asf", matchers.Asf)
	class     = newNode("application/x-java-applet; charset=binary", "class", matchers.Class)
	swf       = newNode("application/x-shockwave-flash", "swf", matchers.Swf)
	crx       = newNode("application/x-chrome-extension", "crx", matchers.Crx)
	woff      = newNode("font/woff", "woff", matchers.Woff)
	woff2     = newNode("font/woff2", "woff2", matchers.Woff2)
	wasm      = newNode("application/wasm", "wasm", matchers.Wasm)
	shp       = newNode("application/octet-stream", "shp", matchers.Shp)
	shx       = newNode("application/octet-stream", "shx", matchers.Shx, shp)
	dbf       = newNode("application/x-dbf", "dbf", matchers.Dbf)
	exe       = newNode("application/vnd.microsoft.portable-executable", "exe", matchers.Exe)
	elf       = newNode("application/x-elf", "", matchers.Elf, elfObj, elfExe, elfLib, elfDump)
	elfObj    = newNode("application/x-object", "", matchers.ElfObj)
	elfExe    = newNode("application/x-executable", "", matchers.ElfExe)
	elfLib    = newNode("application/x-sharedlib", "so", matchers.ElfLib)
	elfDump   = newNode("application/x-coredump", "", matchers.ElfDump)
	ar        = newNode("application/x-archive", "a", matchers.Ar)
	dcm       = newNode("application/dicom", "dcm", matchers.Dcm)
)
