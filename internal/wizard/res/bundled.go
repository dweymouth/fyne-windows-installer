package res

import "fyne.io/fyne/v2"

// This is a dummy file representing the bundled metadata for an application.
// It is for building and running the wizard on its own and will be replaced
// with a file generated by `fyne bundle` containing the installed app's metadata.

var ResAppName = &fyne.StaticResource{
	StaticName:    "app-name",
	StaticContent: []byte("Dummy App"),
}

var ResAppVersion = &fyne.StaticResource{
	StaticName:    "app-version",
	StaticContent: []byte("0.0.1"),
}

var ResAppIcon = &fyne.StaticResource{
	StaticName: "app-icon",
	StaticContent: []byte(`<?xml version="1.0" encoding="iso-8859-1"?>
	<!-- Uploaded to: SVG Repo, www.svgrepo.com, Generator: SVG Repo Mixer Tools -->
	<svg fill="#000000" version="1.1" id="Capa_1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" 
		 viewBox="0 0 495.61 495.61" xml:space="preserve">
	<g>
		<g>
			<g>
				<path d="M441.899,0H53.754C24.224,0,0.043,24.202,0.043,53.754v388.102c0,29.574,24.181,53.754,53.711,53.754h388.102
					c29.574,0,53.711-24.224,53.711-53.754V53.754C495.632,24.202,471.429,0,441.899,0z M427.921,44.414
					c13.978,0,25.367,11.346,25.367,25.367s-11.325,25.367-25.367,25.367c-14.043,0-25.367-11.346-25.367-25.367
					S413.922,44.414,427.921,44.414z M348.778,44.393c13.999,0,25.389,11.346,25.389,25.389S362.799,95.17,348.778,95.17
					s-25.389-11.346-25.389-25.389S334.778,44.393,348.778,44.393z M454.971,449.514H40.747V131.56h414.224V449.514z"/>
				<rect x="128.13" y="317.047" width="79.229" height="79.208"/>
				<rect x="291.551" y="317.047" width="79.251" height="79.208"/>
				<rect x="128.13" y="188.291" width="79.229" height="79.186"/>
				<rect x="291.551" y="188.291" width="79.251" height="79.186"/>
			</g>
		</g>
	</g>
	</svg>`),
}