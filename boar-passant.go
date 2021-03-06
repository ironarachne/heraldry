package heraldry

import svg "github.com/ajstarks/svgo"

func renderBoarPassantToSvg(canvas *svg.SVG, hexcode string, lineColor string) {
	pathData0 := "M89.2,195.1c-5.7-2.5-10.6-4-14.8-6.6c-4.5-2.8-8.2-2.7-13.4-1.4c-9.6,2.4-19.8,2.1-29.2-2c-2.9-1.3-4.8-4-8.1-0.6    c-1,1-3.7,1-5.2-0.8c-1.5-1.8-1.2-3.7-0.4-5.7c1.3-3.2,4.3-4.8,6.8-6.9c-7.2-1-11.1-5.2-10.8-11.4c0.3-4.8,5.5-8.5,9.5-6    c2.2,1.4,3.8,1,5.6,0.4c5.1-1.5,10.2-3.2,15.2-4.9c9.1-3,12.6-10.2,14.7-18.9c2.1-8.7,5.9-16.3,15.3-19.6c2-0.7,4.2-1.6,5.6-3.1    c3.2-3.5,3.9-0.6,4.8,1.6c1.7,4.1,2.1,8.4,1.5,13c6.1-2.8,11.3-6.4,16.3-10.5c1.1-0.9,2.5-2.3,3.9-1.6c1.2,0.6,1.8,2.3,1.5,3.9    c-0.8,3.9,1,4.5,4.5,4.3c14-0.9,27.8,1.1,41.6,3.2c19.2,2.9,37.3,9,55.2,16.1c12,4.7,22.5,11.8,32.3,20c4.5,3.8,5.9,9.8,8.6,14.8    c2.5,4.8,7,6.9,11.1,9.6c1.5,1,2.5-0.6,4.3-0.9c-4.5-3-5.9-6.9-5.5-11.9c0.3-3.6,1.4-6.2,4.6-7.9c3.4-1.8,6.1,0.1,8.4,1.9    c5.2,4.1,6.3,11.4,2.6,17c-1,1.6,0.2,2.3,0.5,3.3c1.9,4.9,4.7,9.9,3.7,15c-1.1,5.3-0.1,11.9-6.6,15.2c-1.2,0.6-2.2,0.9-3.4,1    c-1,0.1-1.6-0.5-2-1.2c-0.6-0.9,0.3-1.5,0.9-2c2.9-2.4,2.3-3.8-0.5-6.1c-5-4.1-5-7.2-1.3-12.9c1.3-2,3.2-3.8,2.1-6.6    c-7.8,3.9-10.2,5.5-22.4-4.3c5.2,11.4,3.9,23.6,9.7,33.9c2.7,4.8,7,8.7,8.3,14.6c0.6,2.7,0.6,5.4,0.8,8c0.4,5.6,3.1,10.6,3.1,16.2    c0,5.1-3.4,7.7-8.5,6.6c-1.2-0.3-1.3,0.8-1.5,1.7c-1.3,4.9-4.8,7.6-9.6,8.5c-2.4,0.4-4.8-1-7.6-0.1c-2.1,0.7-3.3-2.5-1.7-4.6    c3-3.7,5-7.9,8.6-11.3c3.9-3.7,3.4-8.6,1.9-13.3c-2.9-9.3-7.4-17.7-14.8-24.1c-4.4-3.8-9.7-6.6-14.6-9.9c-1.4-1-2.9-1.9-4.4-3    c0,1.9-0.3,3.5,0.6,5.3c3.7,7.7,0,14.3-3.9,20.5c-4.9,7.8-10,15.5-12.8,24.4c-1.3,4.1-5.3,5.8-9.5,4.2c-1.7-0.6-1,0.7-1.4,1.1    c-2.7,2.7-5.4,5.3-9.1,6.8c-4.2,1.6-6.5,1-8-3.2c-0.6-1.7-1.2-0.6-2.1-0.4c-1.6,0.4-3.5,0.8-4.4-1c-0.9-1.8,0.8-2.7,1.8-3.7    c2.7-2.9,5.7-5.4,9.7-6.3c3-0.7,4.5-3.1,6.3-5.3c4.5-5.6,8.5-11.6,10.9-18.2c2.9-8,1-16-3.3-23.2c-1.4-2.4-2.2-6.2-4.7-6.6    c-3.6-0.5-7.5,0.5-11.1,1.6c-1.8,0.5-3.1,2.6-5.1,2.9c-1.3,0.2-3.2,1-3.6-0.5c-1-3.5-3.3-2.5-5.2-1.6c-1.5,0.6-2.9,0.8-4.4,0.7    c-7.7-0.3-15.1,0.7-22.4,3.3c-0.9,0.3-1.5,0.4-1.4,1.5c0.1,4.1-1.3,7.9-3.2,11.4c-0.6,1.1,0.5,1.9,0.8,2.6    c1.4,2.7,1.6,4.5-2.2,4.6c-1.2,0-1.3,0.7-1.3,1.6c0.3,3.7-0.9,7.4-0.7,11c0.2,5.5-1.2,10.9-0.9,16.4c0.1,2.9-2.5,3.6-3.9,5.2    c-1.4,1.5-2.7,3.8-5.3,1.2c-0.9-0.9-0.5,0.7-0.7,1.1c-3.7,7.2-8.4,8.1-15.1,3c-0.3-0.2-0.7-0.3-1.1-0.4c-4.5-1.3-4.5-1.7-1.4-5.4    c2.8-3.3,6.7-5.3,9.6-8.6c1.6-1.9,2.3-3.8,2.2-6.3c-0.1-6.1,0.1-12.2-0.1-18.3c-0.1-5.6,1.7-11.2,0.5-17.3    c-8.1,3.9-17.1,4.4-25.3,7.7c-5.3,2.2-6.7,7.8-8.3,12.8c-0.6,1.7-1.3,3.5-1,5.5c0.3,2,0.5,4.4-1.5,5.5c-2.1,1.2-3.3-1.3-5-2.2    c0,1.7-0.3,3.2,0.2,4.7c1.3,4.6-0.4,8.1-4.6,9.9c-0.6,0.2-1.7,0.4-1.9,0.1c-1.7-2.3-3.5-2.4-5.7-0.7c-0.7,0.6-1.6-0.3-2-1    c-0.3-0.6-0.4-1.4-0.2-2c2.7-7,2.7-14.9,7.3-21.3c4.1-5.8,7.8-11.8,9.3-18.9c0.9-4.4,3.7-7.3,7.8-9    C82.3,198.5,85.1,197,89.2,195.1z"
	pathData1 := "M78.4,186.7c-2.3-1.3-4.7-2.3-5.7-5.2c-0.7-1.9-1.5-1.7-2.9,0.2c-0.1,0.2-0.4,0.4-0.6,0.5    c-7,0.1-13.1,5.2-20.4,3.5c-4.2-1-8.6-1-12.8-1.8c-1.9-0.4-4.7-1.4-4-3.6c0.5-1.5,3.3-1.8,5.4-0.5c1.9,1.2,3.9,1.6,6-0.3    c1-0.9,2.7-1.2,4.2-1.2c5.9-0.2,11.9,0.4,18.1-1.2c-0.7,2.4-3,2.8-3.2,5c2.9-0.8,5.4-1.9,5.8-5.1c0.4-2.8-1-4.9-3.5-6.5    c-0.6,1.7,1.3,2.6,0.9,4.1c-1.4,0.7-1.8-0.7-2.6-1.3c-3.4-2.7-5.4-7.1-10.2-7.9c-0.7-0.1-1.1-1-1.2-2c-0.6-3.5-3.6-5-6.1-6.7    c-1.1-0.7-3-0.9-4.2,0.5c-1.4,1.8,0.7,2.4,1.5,3.1c2.7,2.4,3.1,4.3-0.5,6c-0.3-0.3-0.5-0.5-0.6-0.7c-0.5-1.4,0.8-3.4-0.7-4.1    c-2.2-1-2.6,1.4-3.4,2.7c-2.1,3.6-5.2,5.2-9.2,5.9c0.5-3.6-1.6-6.6-2.4-10c-0.3-1.5-0.6-2.4,1.7-3.1c7.7-2.2,15.2-4.9,22.7-7.6    c4-1.4,6.5-5.1,8.7-8.1c5.7-7.6,13.1-12,21.8-14.9c1.2-0.4,2.4-0.9,3.9-1c-5.6,2.8-9.6,7.3-13.4,12.8c3.2-1,4.6-3.4,6.3-5.2    c1.8-1.9,2.8-4.8,6.2-4.4c0.4,3.5-1.9,7-0.3,11.1c2.2-5.8,1.5-12.4,6.7-16.9c4.6-4,10.7-5.7,15.5-11.3c-0.8,4,1.2,7.4-1.3,9.6    c-4.3,3.6-6.6,9.6-13.2,11.4c3.7-5.7,9.6-8.4,13.4-14.6c-5.9,2.4-9.4,6.2-13.3,9.6c-4.4,3.8-5.3,8.5-5.1,13.7    c0.9,0.1,1.8-0.8,1.6-1.3c-2.1-4,2.3-3.8,3.6-5.6c0.5-0.8,0.9,0,1.2,0.3c0.6,0.6-0.1,1-0.4,1.3c-1,0.9-2.9,0.6-3.3,2.4    c1.6,1.4,2-1,3.1-0.8c0.9,1.4-1.3,1.6-1.1,2.8c2.6,0.1,7-3.3,7.1-5.4c0-0.6-0.6-1.5-1-1.6c-1.8-0.4-1-1.1-0.3-1.7    c0.3-0.3,0.6-0.6,0.9-0.9c1.9-2,3.8-4.1,6-6.5c0.9,4.1-6.9,6.1-1.3,10.2c-3,2.3-4.6,5.8-8.2,7.3c-1.4,0.6-0.3,2.4-1.5,4.5    c3.7-2.3,6.4-4.1,9.1-5.8c-1.8,3.6-4.8,5.7-8.5,7.1c5.5,0.9,12.6-2.5,15.7-7.4c-1.3-0.3-2,2.2-3,1c-0.7-0.8-0.7-2.4,0.4-3.3    c1.4-1.2-0.2-2.4,0-3.7c0.9-5.2,2.8-8.8,7.1-12.5c-2.7-0.8-3.1,2.3-5.1,2.8c-1.2-3-0.1-3.9,3-3.9c10.8,0,21.5,1.3,32.2,2.2    c17.8,1.5,34.6,6.5,51.1,12.2c11.9,4.1,23.1,9.9,34.2,16.4c9.9,5.8,14.4,14.4,19.6,23.2c2.7,4.6,7.4,6.6,11.7,9.1    c1.2,0.7,3.5,0.8,4.9-1c0.7-0.9,1.5-0.5,2,0.2c0.9,1.4-0.8,0.8-1.1,1.3c-2.9,3.6-8.1,2-14.2-1.6c-7.6-4.3-8.2-12.7-12.7-18.9    c-1.6-2.1-3.5-3.8-5.1-5.8c-0.8-1-1.5-0.9-2.4,0c-0.8,0.9-0.3,1.4,0.3,2.1c9.8,10.1,13,23.5,15.5,36.4    c2.1,10.9,5.5,20.5,11.8,29.4c3.5,4.9,3.7,11.1,4.3,17c0.2,2.5,0.4,4.8,1.6,7c2.3,4.3,0.1,6.8-3.3,8.9c-2.9-3.1-0.6-5.6,1.6-8.9    c-4.3,1.6-4.7,4.6-5.7,7.3c-1.8,4.9-3.4,10.2-10.2,11.5c1.1-2.2,2-3.9,3-5.6c0.2-0.3,0.5-0.8,0.8-0.9c1-0.3,2,2.6,2.6,0.4    c0.5-1.8-2-2.1-3.3-2.9c-0.6-0.4-1.4-0.7-2.1-1.5c6.9-4.7,5.5-11.1,3.8-17.9c-2.5-10-7.6-18.4-15.3-25.1c-3.5-3-7.4-5.3-11.2-7.9    c-3.9-2.7-7.6-5.6-10.3-9.7c-0.6,0.2,0.2,1.4-0.6,2c-2.2-2.9-6.3-4.2-7.1-8.4c-0.2-1-1.6-1.4-2.1-1.1c-0.7,0.4-0.4,1.5,0.2,2.2    c0.3,0.3,0.6,0.6,0.9,0.9c0.5,0.6,1.6,1.1,0.8,2c-0.7,0.8-1.7,0.4-2.5-0.4c-3.2-2.9-6.4-5.7-10.8-6.7c-2.7-0.6-3.4-3.9-6-4.8    c-1.1,1,1.5,1.8-0.2,3.1c-1.9-5.2-6.5-7.6-10.9-11.2c0.7,2.5,0.6,4.3,3.1,5.8c3.1,1.8,4.2,5.9,7.7,7.6c1,0.5,1.8,1.2,3,1    c1-0.1,2.1,0.1,2.5,1c0.6,1.4-2.9-0.4-1.7,2c0.2,0.3-1.3,0-1.8-0.6c-3.2-3.1-8.7-3.3-10.4-8c-1.3-3.5-4.3-4-7.2-5.1    c3.2,2.6,5.8,5.5,6.7,9.6c-3.4,1.4-5.3-2.6-8.3-3.2c1.7,2.7,3.9,4.6,7,5.1c0.9,0.2,1.8,0.9,2.6,1.6c0.4,0.3,0.4,0.9,0.6,1.4    c-0.5,0.1-1.1,0.4-1.6,0.2c-2-0.8-4-1.5-5.9-2.5c-0.9-0.4-1.7-1-2.7-0.7c-7.6,2-14.1-3.6-21.4-3.5c-1.8,0-2.8-2.6-5-2.3    c0,2.8,2.3,3.5,4.2,4c7.8,2.3,15.4,5.7,23.6,5.8c3.5,0.1,1.6,2.4,1.2,2.9c-0.9,1.2-2.4,2.1-3.8,2.9c-0.3,0.2-1.5-0.1-1.5-0.3    c-0.2-5.5-4.4-3.9-7.2-3.6c-3.6,0.3-3.6,0.7-5.9-2.1c-0.6-0.6-1.4-1-1.8-0.4c-1.3,1.9-3.2,1.3-5.1,1.5c0.9,1.5,3,1.2,3.7,2.8    c-0.3,1.6-1.4-0.3-2.4,0.3c-1.9,1.2-4.6,2.9-6.9,0.1c-1.4-1.6-2.9-0.9-4.5,0c1.2,1.4,2.2,0.5,3.6,0c-0.4,0.9-0.5,1.4-0.7,1.5    c-1.7,0.4-3.4,0.7-5.1,1c-0.9,0.2-2,1.3-2-0.9c0-1-1.4-1.9-0.5-3c0.9-1.1,1.5,0.8,2.5,0.5c0.4-0.1,0.8-0.1,1.2-0.2    c-1-2.5-4.3-1.4-5.5-3.6c1.2-0.9,2.1,0,3,0.5c1,0.5,2.4,1.7,3-0.1c0.4-1.3-1.3-1.1-2-1.6c-2.2-1.5-4.3-3.2-5.6-6.5    c4.2,1.2,6.3,4.1,9.3,5.6c-2-4.5-6.8-5.6-9.7-8.8c0.2-0.1,0.9-0.3,1.6-0.5c-1.7-1.2-3.3-2.4-6-2.7c2,6.8,3,13.6,5.3,20.2    c-1.4-3-5.1-2.5-7-4.8c3.4,3.8,5.5,8.2,7.4,13.1c-4.3-0.7-4.9-4-6.3-6.6c-0.4-0.7-1.2-1.9-1.4-1.8c-1.1,0.4-2.6,0.5-2.7,2.1    c-0.1,2,0.1,0.8,1.1,0.4c1.4-0.6,8,9.5,7.5,11.2c-0.1,0.2-0.4,0.4-0.8,0.7c-1-1.7-1.7-3.6-3-5c-1.1-1.2-1-4-4-3    c2.7,5.3,6.8,9.9,6.8,16.7c-3.2-2.4-2.9-6.4-5.4-9.1c3.5,13.3,0.1,26.5,0.1,40.1c-3.3-0.8-2-4.8-4.6-6.1c-1.1,1.5,1.6,5.2-2.9,5.1    c2.7-0.1,3.3,2.4,5.2,3c0.9,0.3,0.1,0.8-0.1,1.1c-0.6,0.7-1.2,1.6-2.4,1.3c-0.3-0.1-1,0.1-0.6-0.6c1.2-2.4-1.1-3.5-2.1-3.3    c-1.8,0.3-2.2,2.4-2.4,4.2c-0.4,4.3-3.1,6.7-6.7,6.4c-1.7-0.1-2.7-1.1-3.1-2.3c-0.5-1.5,1.4-1.4,2.2-1.8c1.8-0.9,3.3-3.4,5.8-1.5    c0.6,0.5,0.9,0.2,1-0.5c0-0.7,0-1.3-1-1.4c-3.3-0.4-6.2,0.6-8.7,2.6c-0.9,0.8-1.6,2-3.1,1.4c4.4-6.8,4.4-6.8,9-5    c0.3-1.5-1.1-1.4-1.9-2c5.3-3.4,6.8-7.7,6-14.2c-1.3-11.3,0.2-22.7,1.4-34c0.2-2.3,1.6-0.4,2.5-0.8c-3.1-6.3-11.5-9.7-9.1-19.4    c-2.4,2.1-1.1,3.7-2.2,4.9c-2.8-3.5-3.3-7.3-1.7-11.4c-3.8,2.8-2.8,6.6-1.9,10.8c-1.3-1.3-1.6-2.8-3.1-3.2c-1,2.3,0.2,4.2,1.5,6.6    c-4.1-0.9-3-5.8-6.5-6.4c-1.2,1.4,0.5,2,1,3.1c0.5,1,1.6,2,0.9,3.6c-2.6-1.4-5.7-2.1-4.7-6.7c-5.2,6.3-8.5,0.2-12.4-1.2    c1.5-2.6-1.5-3.4-1.8-4.6C77.2,183.1,79.2,184.4,78.4,186.7z"
	pathData2 := "M198.9,218.9c0.9,0.6,1.4,1,2.2,1.6c0.3-2.6-4.1-2.6-1.7-6c3,3.6,6.1,7.1,9.1,10.7c0.5-0.5,1-0.9,1.4-1.4    c-4.3-4.2-8.6-8.5-13.6-13.5c7.2,1.7,9.5,7.9,14.8,11.8c-4.3-9.4-13.7-12.4-19.1-19.8c3.2,0.1,4.9,3.1,8.1,3.2    c-2.2-3-5.7-3.8-8-6.4c3.7-2.5,5.7,1,7.8,2.4c3.1,2.1,5.6,5.2,8.2,7.9c0.9,0.9,1.7,1.3,3.2,0.7c-4.3-3.8-7.6-8.4-12.2-11.8    c3.3,0.2,5.2,2.8,8,4.8c-0.3-2.9-3.2-3.6-4.1-6.2c4.8,1.4,6.4,6,11,7.8c-3.3,1.5-3.7-2.1-6.6-1.3c3.9,3.5,7.5,6.8,7.8,12.1    c-0.9-0.8-1.6-2-2.5-2.2c-1-0.2,0.1-2.6-1.6-1.3c-1.2,0.9,0.2,1.8,0.5,2.4c0.7,1.4,1.9,2.6,2.9,3.9c0.3,0.4,1.3,0.7,0.4,1.6    c-1.1,1.1-1.4-0.1-1.8-0.5c-2.4-2.2-3.8-5.2-5.9-7.7c-1.7-2-3.4-3.6-6-5.2c0.9,2.9,3,4.7,4.6,6.2c2.7,2.7,4.3,6.1,7.3,8.5    c0.9,0.7-1.3,6.7-2.7,7.3c-2.2,0.9-2.5-1.2-3.3-2.5c-0.7-1.1-1.6-2-2.7-2.8c0.4,3.4,3.1,5.4,5.2,7.7c-1,1.1-1.7,2.9-4.1,1.5    c1.6,2.2,0.9,3.9-0.4,5.7c-2.7,3.7-4.4,7.9-6.2,12.1c-1.1,2.4-1.5,5.3-4.6,6.2c-0.6,0.2-1,1.1-1.7,0.4c-0.5-0.5-0.5-1.5-0.1-1.9    c1.6-1.3,1.2-3.8,2.8-5.1c-2.8,0.8-4.7,2.9-6,5.1c-2.6,4.2-6.4,6.8-10.1,9.6c-1,0.8-2.7,2-3.9,0.1c-1.3-2,1.4-3.8,1.6-3.8    c2.8-0.2,4.1-3.1,7.5-3c-2.9-1.5-5.8,1.8-7.7-1.4c0,1.9-1.9,1.9-2.9,2.9c-1.2,1.3-2.6,2.3-4.3,1.9c2.8-3.9,6.8-6,11.4-7.3    c3.6-1.1,4.6-4.7,6.6-7.4c4.2-5.9,8.4-11.7,10.7-18.7C199,224.9,198.2,222.1,198.9,218.9z"
	pathData3 := "M86.2,207.7c3.8-0.8,3.8-4.5,5-7c2.2-4.6,4.9-0.8,7.1-0.7c2.2,0,1.4,2.3,0.6,3.9c-0.6,1.2-1,2.6-1.7,4.4    c3.8-1.6,3.1-6.1,6.4-7.5c0.2,1,0.3,2,0.6,3c0.1,0.6,1.6,1.6,0.7,1.5c-1.7-0.1-1.1,3.8-3.3,1.7c-1.1-1-0.8,0.3-1.1,0.7    c-1.4,1.7-3.5,1.8-5.5,2.5c0.7-2.7,1.5-5.5,2.3-8.4c-0.2-0.1-0.3-0.2-0.5-0.3c-1.8,3.2-3.6,6.4-5.4,9.7c-1.7-1.1,0.1-2.2-1-4.4    c-0.9,5.7-4.8,6.6-8.5,8.1c-1.6,0.6-3,2.2-3.8,4.2c-2.1,5.7-6,11-4.8,17.6c-5.3-0.4-2.3-5.5-4.8-7.4c-0.8,4.5-1.8,9.1-2.5,13.7    c-0.3,1.8,2.7,4-0.6,5.7c-0.9,0.4-1.4,1.3-2.4,0.3c-0.7-0.8-0.8-1.4,0-2.1c3-2.8,1.2-4.3-2.2-5.9c3.3,4.9-1.6,5.3-3.3,7.5    c0.3-3.5,2.7-6.3,3.2-9.7c1-5.9,5-10.5,7.7-15.6c2.1-4,4-8.1,5.5-12.4c1.9-5.2,5.1-8.6,10.4-10.2c0.5-0.2,1-0.6,1.5-0.5    c1.6,0.1,2.4-3.3,4-1.4c1.3,1.7-2.1,1.4-2.1,2.7c0,4.8-5.6,5.1-6.9,9.3c2.2-0.5,3-2.4,4.4-3.6c1.5-1.2,1.1-0.1,1,0.7    c-0.3,1.3-0.4,2.6-1.6,3.4C85.9,210.4,86.2,209.1,86.2,207.7z"
	pathData4 := "M278.3,199.3c-0.7-3.5-3.1-6.4-3.5-10.1c-0.8,4.7,1.7,8.8,2,13.3c-1.7-3.9-4-7.7-3-12.5    c-1.7,2.1-2.1,3-1.7,4.2c0.8,2.4,1.5,4.8,2.5,7.2c2.7,6.4,2.6,6.9-3.5,11.5c3.8-4.4,1.2-7.1-2-10c-2.6-2.3-2.7-5.3-0.8-8.1    c4.4-6.6,4.3-10.1-1.5-15.7c-2.8-2.7-4.3-5.3-3.9-9.3c0.3-3.1,1.6-5.5,4.1-6.5c2.9-1.1,5.3,0.8,7,3c2.1,2.7,2.5,5.9,2.3,9.3    c-0.1,1.7-1.3,3.2-2.2,3c-1.5-0.3-1.7-2.3-1-3.4c1.6-2.8,0.3-4.6-1.6-6.4c-1.6-1.5-3.1-2.3-4.9-0.5c-1.8,1.8-1.7,3.6-0.1,5.5    c3.2,4,7.3,7.3,8.8,12.4C276.3,190.4,278.4,194.5,278.3,199.3z"
	pathData5 := "M78.3,118.4c1.5-1.1,3.1-2.1,4.8-3.3c0.2,2.5-3.4,2.9-2.5,5.7c0.9,0.7,1.2-1.4,2.2-1.4    c0.4,3.6,0.3,3.3-2.5,4.9c-6.2,3.3-13.9,3.9-18.4,11.2c0.3-12.9,6.9-21.5,19.1-25.5c0.4-0.1,0.8-1.1,1.2-0.2    c0.3,0.7,0.5,1.5-0.7,1.9c-1.6,0.6-3,1.4-4.3,3.3c2.5-0.1,3.4-2.8,5.6-2.3c-0.2,3-5.1,2.3-4.4,5.9c-0.3,0.1-0.5,0.1-0.4,0    C78,118.4,78.1,118.4,78.3,118.4z"
	pathData6 := "M35.4,170.1c1.8,4.2,0.3,5.8-4,6.6c-3,0.6-5.6,2.9-8.3,4.6c-0.8,0.5-1.2,1.6-2.4,0.6c-1-0.9-1-1.8-0.4-2.9    c1.8-3.3,4.3-5.9,8.2-6.5C31,172,33.1,170.9,35.4,170.1z"
	pathData7 := "M42.7,157.7c4.6,1.9,7.2,3.9,7.2,8.2c0,2.8-1.3,4.5-3.8,5.5c-2.6,1-3.1-1.3-4.4-2.8c2.4-0.6,4.4-1.4,4.9-4.1    C47.1,161.5,44.7,160.4,42.7,157.7z"
	pathData8 := "M18.9,164.2c0.8-0.9,2.3,1,2.6-0.2c0.5-1.8-1.5-3-1.7-4.8c-0.1-1.4-3.3-0.4-1.9-2.7c0.5-0.8,1.3-0.8,2.1-0.9    c2.1-0.2,2.5,1.5,3,2.8c1.3,3.3,2.4,6.7,3.5,10c-3.6,0.5-5.4-2.2-7.5-4.3L18.9,164.2z"
	pathData9 := "M58.5,173.2c-1.3,0-2.6,0-3.9,0c0,0.2,0,0.5,0,0.7c1.9,0,3.8,0,5.7,0c-5-0.5-10.3,4.2-15.8-0.5    c3.5-0.1,5.8-1,7.1-3.8c1-2.1,2.5-1.3,3.9-0.1c-0.7,1-2.2-0.3-2.7,1c1.5,0.1,3,0.2,4.4,0.4c0,0.1,0,0.3,0,0.4    c-1.2,0.1-2.3,0.2-3.5,0.3c0,0.2,0,0.3,0,0.5c1.7,0.1,3.4,0.2,5,0.3C58.6,172.7,58.5,172.9,58.5,173.2z"
	pathData10 := "M39.8,164.5c-0.1,8.3,2.1,8.8,4.9,12.6c-2.1,0-4,0-5.9,0C40.2,172.9,34.2,168.1,39.8,164.5z"
	pathData11 := "M241.9,268.6c0.9-3,2.4-5.2,4.4-7.3c1-1.1,1.7-0.3,2.4,0c1.3,0.6-0.1,1.6-0.2,1.8    C246.7,265.2,245.1,267.6,241.9,268.6z"
	pathData12 := "M270.8,175.3c-1.4-1.3-3.9-1.9-2.9-4.4c0.5-1.3,1.3-0.6,2.2-0.3C273.6,171.7,270.3,173.7,270.8,175.3z"
	pathData13 := "M19,164.1c-1.2-1-2.8-1.8-2.3-3.9c0.1-0.2,0.8-0.5,1-0.3c1.8,1,1.4,2.7,1.1,4.3C18.9,164.2,19,164.1,19,164.1    z"
	pathData14 := "M73.4,148.3c-1.5,0.9-2.4,1.4-4.1,2.4c3,0,5.1,0,7.1,0c0.2,0.4,0.5,0.8,0.7,1.2c-7-0.5-4.2,9.3-11.6,8.7    c2.2,1.7,4,3.1,6.1,4.7c-8.1-1.2-11.7-5.2-10.7-10.8c0.4-2.1-1-1-1.6-1.4c-1-0.5-2.9-0.3-2.7-1.9c0.3-2.1,1-4.3,3-5.4    c1.2-0.7,2.6,0.2,2.7,1.6c0.1,1.5,2.6-0.1,1.7,1.4c-0.4,0.7-1.9,1.6-3.2,0.6c-0.7-0.6-1.5-0.7-2.3-0.2c-0.2,0.1-0.1,0.7-0.2,1.1    c0.8,0.1,1.6,0.3,2.4,0.3c2.8,0.1,5.6-0.2,8-1.9C70.1,147.8,71.4,148.3,73.4,148.3z"
	pathData15 := "M73.1,171.4c3.6-0.9,6.7,0.5,9.8,1.7c4.9,1.9,9.3,5.4,15,5c1.2-0.1,1.8,1.1,2.4,2c0.1,0.2-0.4,1.7-1.5,0.8    c-0.4-0.4-0.8-0.8-1.2-1.1c-0.3-0.2-0.8-0.5-0.9-0.4c-0.6,0.6,0.6,2.6-1.6,1.9c-0.8-0.3-1-1.7-2.7-1.2c1.4,1.6,2.3,3.4,2.8,5.6    C88.9,179.1,81.6,174.4,73.1,171.4z"
	pathData16 := "M85.5,187c-3.4,0.8-3.4-2.2-4.7-3.5c-1.9-2-2.9-4.7-6-6c-2.6-1.1-3.9-0.6-5.8,2c0.4-3,4-5.5-1.7-7.5    c0.8-0.2,1.6-0.5,2.3-0.7c1.2,1.3,2.1,3,3.6,3.9C76.6,177.5,83.1,183.5,85.5,187z"
	pathData17 := "M109.6,175.8c-2.1-2.1-4.3,0.3-6.4,0.3c-4.8,0-7.7-3.2-11.2-5c-1.3-0.7-3.8-1.1-4.4-3.3l-0.1,0.1c2.5-2,5.3-1.4,6.8,0.5    c1.4,1.8,3.3,2.6,4.7,4.2c1.7,1.9,3.8,3.2,6.7,1c2-1.6,3.4,0.1,3.8,2.3L109.6,175.8z"
	pathData18 := "M94.4,160c1.2-1.4,1.9-2.2,2.8-3.1c-2.7-1.1-4.9,0.6-7.3,1c1.6-2.6,3.1-3,12.9-3.4c-1.8,1.5-5.4,1.5-3.5,4.9    c1.6-2.6,4.8,0.1,7.1-3.1c-1.1,3.1-1.1,3-4.5,4.3C99.4,161.5,97.6,158.8,94.4,160z"
	pathData19 := "M33.5,180.7c9,2.8,17.9,0.8,26.8,0.3c-0.8,1.8-1,1-3.5,1.5c-4.2,0.8-8.7-0.5-12.5,0.8C39.8,184.8,37,182,33.5,180.7z"
	pathData20 := "M140.9,194.8c2.2,0,3.7,0,5.2,0c0.6,0,1.1,0.3,1,0.9c-0.1,0.3-0.7,0.7-0.8,0.6c-1.6-1.2-3.5-0.5-5.1-0.9    c-0.8-0.2-1.6-0.9-1.3-1.7c0.5-1.4,1.3-2.6,2.1-3.9c0.1-0.1,1.9-0.9,1.6,1.1c-0.1,0.6,1.3,0.4,2.2,0c1.1-0.5,2.6-1.2,3.3,0.8    c0.3,0.9,2.8,0.5,1.7,2.2c-1.1,1.8-1.9-0.4-2.9-0.6C145.7,193,143.4,192.3,140.9,194.8z"
	pathData21 := "M98.1,152.3c-4.4,3.2-9.5-2.4-13.8,1.9c2.7-3.7,6.5-3.2,10-3.6c0.1,0.6,0,0.9-0.5,0.9c-0.4,0-0.7,0.2-1.3,0.3    C94.5,153.2,96.3,151.4,98.1,152.3z"
	pathData22 := "M96.9,176.7c-4.9-1.4-8.4-4-12.5-4.7C87.1,169.6,91.8,171.1,96.9,176.7z"
	pathData23 := "M128.2,130.3c3.4-3.6,7.2-5.8,12.2-5.7C135.3,124.4,133.4,130.8,128.2,130.3z"
	pathData24 := "M85.5,189.3c4.7-2-1.6-6.4,2.3-8.5c0.5,1.6,2,2.4,2.8,3.7c1.6,2.3-1.5,1.3-1.6,2C89,188.7,87.3,188.8,85.5,189.3z"
	pathData25 := "M65.2,138c3-3.8,6.5-6.5,11-7.9C72.2,132.3,70,137.1,65.2,138z"
	pathData26 := "M130.5,122.6c-2.4,3-5.5,5.2-7.5,8.3C123.2,127,127,123,130.5,122.6z"
	pathData27 := "M72.8,167.6c3.8-0.2,7.5,0.5,12,2.8C79.8,170.8,76.4,168.7,72.8,167.6z"
	pathData28 := "M126,121.7c-2.8,3.5-6.1,6.4-9,9.8C119.2,127.6,121.9,124.1,126,121.7z"
	pathData29 := "M113.6,143.4c-1.2,3.3-3.9,5-6.9,6.2C108.7,147.2,110.6,144.7,113.6,143.4z"
	pathData30 := "M78.4,186.7c-0.9-3-3.2-5.1-4.9-7.6c3.1,1.1,4.6,4.1,7.2,5.9c-0.9,0.6-1.6,1.1-2.4,1.6C78.3,186.6,78.4,186.7,78.4,186.7z    "
	pathData31 := "M173.7,182.9c-3.4,0.9-5.2-1.3-7.1-3.2c-0.3-0.3-0.7-1,0-1.4c0.3-0.2,1-0.1,1,0C168.6,181.2,172.1,180.4,173.7,182.9z"
	pathData32 := "M244.1,195.4c-3-1.9-6.1-3.7-9.3-5.4C238.5,190.6,241.8,192.1,244.1,195.4z"
	pathData33 := "M87.6,167.8c-3.5,1-5.2-2.1-8.3-3.5c4.2-0.9,5.5,2.7,8.2,3.6C87.5,167.9,87.6,167.8,87.6,167.8z"
	pathData34 := "M98.7,170c1.9-0.1,3.4-0.5,4.8,0.3c-2,0.3,0.5,2.7-1,2.7C101.1,173,99.5,172.3,98.7,170z"
	pathData35 := "M115,199.1c0.8,2.3,4.3,3.2,3.7,7C116.1,204,115.5,201.6,115,199.1z"
	pathData36 := "M104.8,167.4c-2.7,2.3-5.4,1.3-7.9,1.2C99.1,167.2,101.7,167.2,104.8,167.4z"
	pathData37 := "M100.6,165.2c-1.3,1.6-2.8,1.8-4.3,0.6c-0.7-0.6-1-1.8-0.5-2.1C97.8,162.5,98.8,164.9,100.6,165.2z"
	pathData38 := "M110.4,128.1c2.1-2.2,3.8-4.9,7.5-4.5c-2.6,1.3-4.3,4.2-7.6,4.4L110.4,128.1z"
	pathData39 := "M233.1,174.2c2.8,0.9,4.2,3.3,7,5.9C236.3,178.9,235.1,176.1,233.1,174.2z"
	pathData40 := "M111.3,131.4c0.4-2.9,3.3-3,4.9-4.5c-1.1,2.1-2,4.5-5,4.4L111.3,131.4z"
	pathData41 := "M116.6,145.8c-0.4,2.7-2.8,3.4-4.9,5.6C112.1,147.8,115.1,147.5,116.6,145.8z"
	pathData42 := "M45.4,180.2c3-0.1,5.9-2,10.1-0.6C51.4,179.5,48.5,181.2,45.4,180.2z"
	pathData43 := "M80.2,136.6c-1.2-3,0.9-4.5,1.7-6.8C83.1,132.8,81,134.3,80.2,136.6z"
	pathData44 := "M239.5,201.6c-1.4-1.9-2.7-3.8-3.5-4.8C237.3,197.2,238.4,199.5,239.5,201.6z"
	pathData45 := "M101,162.9c1.4,1.4,3.8,0.6,4.4,3.1C104,165,101.1,166.9,101,162.9z"
	pathData46 := "M89.9,164.4c3.3-2.2,3.8,1.5,5.3,2.9C93.4,166.3,91.7,165.4,89.9,164.4z"
	pathData47 := "M236.5,173.9c-1.5-1.5-2.9-3.1-4.6-4.9C234.7,169.5,235.6,171.6,236.5,173.9z"
	pathData48 := "M111.2,131.3c0.6,2.5-1.1,3.3-3.3,4.7c0.6-2.6,2.2-3.5,3.4-4.6C111.3,131.4,111.2,131.3,111.2,131.3z"
	pathData49 := "M112,141.3c-0.3,2.2-2.3,2.9-3.7,4c0.6-2,1.7-3.4,3.8-3.9L112,141.3z"
	pathData50 := "M118,150.8c-0.2-2.4,1.3-3,2.9-3.6C121,149.6,117.9,148.6,118,150.8z"
	pathData51 := "M112.1,141.4c1.1-1.3,2-2.9,4-3c-1.2,1.1-1.6,3.5-4.1,2.9C112,141.3,112.1,141.4,112.1,141.4z"
	pathData52 := "M166.4,181.4c-1.6,0.1-3,0-3.7-2.5C164.3,180,165.3,180.7,166.4,181.4z"
	pathData53 := "M103.7,149.3c1-1.7,2.3-2.4,2.5-3.7C106.3,147.7,106.3,147.7,103.7,149.3z"
	pathData54 := "M109.5,175.9c1-0.2,2.1-0.9,2.7,0.5c0.2,0.4,0,1.4-0.3,1.5c-1.9,0.6-1.1-1.8-2.3-2.1C109.6,175.8,109.5,175.9,109.5,175.9    z"
	pathData55 := "M241.2,197.1c-2.1,0.2-2.5-1.1-3.6-2.2C239.7,194.7,240.1,196,241.2,197.1z"
	pathData56 := "M110.3,128c0.8,1.9-0.9,2.5-2,3.7c-0.4-2.1,1.3-2.6,2.2-3.6C110.4,128.1,110.3,128,110.3,128z"
	pathData57 := "M108.4,189.7c-0.6-1.4-0.3-2.3,0.8-3C110.3,187.9,109.7,188.8,108.4,189.7z"
	pathData58 := "M110,138.4c1.3-1.4,1.9-2.1,2.7-3.1C113.3,137.5,112.3,137.9,110,138.4z"
	pathData59 := "M91.1,162.3c-0.9-0.1-1.8-0.1-2.7-0.2c0.2-0.3,0.3-0.8,0.6-1C89.9,160.6,90.4,161.6,91.1,162.3z"
	pathData60 := "M86.2,207.7c1.2,3.9-1.1,4.9-4.9,5.4c1.9-2.1,4.3-2.9,5-5.3C86.3,207.8,86.2,207.7,86.2,207.7z"
	pathData61 := "M80.2,207.9c0.3-1.9,0.7-3.7,2.8-5.3C82.9,205.2,82,206.8,80.2,207.9z"
	pathData62 := "M74.5,125.3c0.5-0.8,0.5-0.9,0.6-1c1.2-1.1,2.4-2.1,3.6-3.1c0,0,0.8,0.7,0.8,0.8C78.7,123.7,77.1,124.3,74.5,125.3z"
	pathData63 := "M67.7,126c1.3-1.6,2-2.6,4.2-3C70.6,124.6,69.9,125.5,67.7,126z"
	pathData64 := "M78.3,118.4c-0.9,0.9-1.4,2.3-3.7,2.4c1.3-1.4,1.7-3,3.8-2.3C78.4,118.5,78.3,118.4,78.3,118.4z"
	pathData65 := "M76.1,116.4c-1.3,1.3-1.8,2.6-3.4,2.6C73,117.6,73.8,116.9,76.1,116.4z"
	pathData66 := "M67.8,158c-5.8,0.2-6,0-4.2-3.9C64.3,155.8,65.4,156.9,67.8,158z"
	pathData67 := "M68.2,153.2c0.7,0.1,1.9-0.3,1.6,1.1c-0.2,0.7-0.8,1.5-1.7,1.3c-0.5-0.1-1.2-0.8-1.4-1.3    C66.3,152.9,67.5,153.3,68.2,153.2z"
	canvas.Path(pathData0, "fill:"+lineColor)
	canvas.Path(pathData1, "fill:"+hexcode)
	canvas.Path(pathData2, "fill:"+hexcode)
	canvas.Path(pathData3, "fill:"+hexcode)
	canvas.Path(pathData4, "fill:"+hexcode)
	canvas.Path(pathData5, "fill:"+hexcode)
	canvas.Path(pathData6, "fill:"+hexcode)
	canvas.Path(pathData7, "fill:"+hexcode)
	canvas.Path(pathData8, "fill:"+hexcode)
	canvas.Path(pathData9, "fill:"+hexcode)
	canvas.Path(pathData10, "fill:"+hexcode)
	canvas.Path(pathData11, "fill:"+hexcode)
	canvas.Path(pathData12, "fill:"+hexcode)
	canvas.Path(pathData13, "fill:"+hexcode)
	canvas.Path(pathData14, "fill:"+lineColor)
	canvas.Path(pathData15, "fill:"+lineColor)
	canvas.Path(pathData16, "fill:"+lineColor)
	canvas.Path(pathData17, "fill:"+lineColor)
	canvas.Path(pathData18, "fill:"+lineColor)
	canvas.Path(pathData19, "fill:"+lineColor)
	canvas.Path(pathData20, "fill:"+lineColor)
	canvas.Path(pathData21, "fill:"+lineColor)
	canvas.Path(pathData22, "fill:"+lineColor)
	canvas.Path(pathData23, "fill:"+lineColor)
	canvas.Path(pathData24, "fill:"+lineColor)
	canvas.Path(pathData25, "fill:"+lineColor)
	canvas.Path(pathData26, "fill:"+lineColor)
	canvas.Path(pathData27, "fill:"+lineColor)
	canvas.Path(pathData28, "fill:"+lineColor)
	canvas.Path(pathData29, "fill:"+lineColor)
	canvas.Path(pathData30, "fill:"+lineColor)
	canvas.Path(pathData31, "fill:"+lineColor)
	canvas.Path(pathData32, "fill:"+lineColor)
	canvas.Path(pathData33, "fill:"+lineColor)
	canvas.Path(pathData34, "fill:"+lineColor)
	canvas.Path(pathData35, "fill:"+lineColor)
	canvas.Path(pathData36, "fill:"+lineColor)
	canvas.Path(pathData37, "fill:"+lineColor)
	canvas.Path(pathData38, "fill:"+lineColor)
	canvas.Path(pathData39, "fill:"+lineColor)
	canvas.Path(pathData40, "fill:"+lineColor)
	canvas.Path(pathData41, "fill:"+lineColor)
	canvas.Path(pathData42, "fill:"+lineColor)
	canvas.Path(pathData43, "fill:"+lineColor)
	canvas.Path(pathData44, "fill:"+lineColor)
	canvas.Path(pathData45, "fill:"+lineColor)
	canvas.Path(pathData46, "fill:"+lineColor)
	canvas.Path(pathData47, "fill:"+lineColor)
	canvas.Path(pathData48, "fill:"+lineColor)
	canvas.Path(pathData49, "fill:"+lineColor)
	canvas.Path(pathData50, "fill:"+lineColor)
	canvas.Path(pathData51, "fill:"+lineColor)
	canvas.Path(pathData52, "fill:"+lineColor)
	canvas.Path(pathData53, "fill:"+lineColor)
	canvas.Path(pathData54, "fill:"+lineColor)
	canvas.Path(pathData55, "fill:"+lineColor)
	canvas.Path(pathData56, "fill:"+lineColor)
	canvas.Path(pathData57, "fill:"+lineColor)
	canvas.Path(pathData58, "fill:"+lineColor)
	canvas.Path(pathData59, "fill:"+lineColor)
	canvas.Path(pathData60, "fill:"+lineColor)
	canvas.Path(pathData61, "fill:"+lineColor)
	canvas.Path(pathData62, "fill:"+lineColor)
	canvas.Path(pathData63, "fill:"+lineColor)
	canvas.Path(pathData64, "fill:"+lineColor)
	canvas.Path(pathData65, "fill:"+lineColor)
	canvas.Path(pathData66, "fill:"+hexcode)
	canvas.Path(pathData67, "fill:"+hexcode)

}