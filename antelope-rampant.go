package heraldry

import svg "github.com/ajstarks/svgo"

func renderAntelopeRampantToSvg(canvas *svg.SVG, hexcode string, lineColor string) {
	pathData0 := "M117.8,118.9c0.6-2.4,1.5-4.5-0.3-6.9c-1.5-1.9,1.7-5.7,4.9-6c1.7-0.2,3.6,0.1,4.6-1.6    c0.2-0.3-0.2-1.3-0.4-1.3c-1.7-0.2-3.1,0.6-4.6,1.2c-5.8,2.3-11.6,4.1-17.4,0.1c-1.9-1.3-3.7-2.8-3.7-5.3c0-2.2,1.8-3,3.5-4    c1.8-1.1,2.9-0.8,3.4,1.1c1.1,4.5,4.7,4.1,7.8,3.5c4-0.8,7.9-2.2,12.1-2.2c1,0,1.9-0.1,1.7-1.4c-0.1-0.8,0.5-1.9-0.6-2.4    c-1-0.4-1.5,0.1-2,1.1c-0.4,0.8-1.2,2.1-2.2,1.7c-1.1-0.4-2.2-1.3-2.2-2.8c0-0.6,0.3-1.4-0.6-1.6s-1.4,0.6-1.7,1.3    c-0.8,2.2-1.8,2-3.2,0.5c-2.7-2.9-7.6-4.6-4.6-10.7c-3.2,1.4-5.9,1.1-7.7,3.2c-1.2,1.4-2.9,2.5-1.5,4.8c0.3,0.5,0.3,1.4-0.4,1.6    c-0.6,0.2-1.2,0.1-1.9-0.7c-2.7-3.4-2-9.8,1.9-12c5.4-3.1,11.2-3.8,17.4-3.1c6.3,0.7,12,4.3,18.6,3.7c2.8-0.2,5.4-0.8,8-1.6    c1.1-0.4,1.8-0.8,1.9-2c0.1-1.5-1.2-0.8-1.8-1.3c-0.9-0.7-2.9-0.7-2.6-2.3c0.3-1.5,1.7-2.2,3.5-2.3c4.4-0.3,7.9,1.3,9.3,5.3    c0.7,2,1.4,1.1,2.1,0.7c0.7-0.3,1.1-1,1.7-1.5c1-0.8,1.7-2.7,3.3-2.1c2.1,0.8-1.2,2.3,0.2,3.4c7.6-6.1,17-7.3,26.2-9    c13-2.4,26.5-2.7,38.8-8.4c2.2-1,3.7-2.9,5.8-4.3c0.9,6-4.3,10.9-14.7,14c1.6,1.3,2.9,2.3,4.2,3.4s1.5,2.4,0.7,3.7    c-1.1,1.7-2.8,0.7-3.6-0.1c-4.5-4.3-10.1-3.7-16.2-3c1.2,1,2.4,2,3.6,3.1c1.3,1.2,1.8,2.5,0.6,4c-1.3,1.6-2.8,0.3-3.7-0.3    c-4.4-3.3-9.5-3.5-14.7-3.7c0,1.4,1.6,1.7,1.6,2.5c-0.3,5.2,3.5,3.6,6.1,3.7c14.9,0.8,29.1-2.1,42.6-8.4c4.8-2.2,8.8-5.6,12.5-9.8    c1.5,6.4-0.9,10.6-8,14.4c-3.1,1.7-6.4,3.1-10,4.8c2.7,1.8,4.9,3.8,6.8,6.1c0.9,1,2,2.9,0.7,4.1c-1,0.9-2.7,0.4-4.1-0.5    c-5.2-3.5-11.1-4.8-17.8-4.4c1.6,2.3,3.9,3.8,4.8,6.4c0.4,1.3,0.5,2.4,0.6,3.7c0.2,1.4-0.9,2.1-1.8,2.7c-1.2,0.8-1.9-0.6-2.4-1.2    c-3.7-4.9-9-7-14.4-9c-1.2-0.4-2.3-0.2-4.1-0.2c3.2,2.9,5.9,5.6,5,9.9c-0.2,0.8,0.5,1.9-0.7,2.3c-1.1,0.4-1.8-0.3-2.4-1.3    c-2.5-3.8-6.2-6.2-9.8-8.6c-1.3-0.9-2.8-1-4.6-0.7c0.2,0.4,0.4,0.6,0.4,0.9c0.1,5,4.1,8,6.2,12c1.7,3.2,1.5,6.1-3.5,5.5    c-3-0.3-6,0.2-9,0.6c-2.6,0.3-2.2,1.7-2,3.1s1.1,2.8,1.2,4.2c0.1,1.6,1.6,2.5,1.8,4c-4.9,0.7-5.9,2.5-3.6,6.2c0.6,1,2,1.8,1.2,2.9    c-0.9,1.3-2.3-0.2-3.5-0.4c-0.4,3.4-0.6,6.6,0.7,9.5c1,2.2,0.4,2.6-1.5,1.8c-2.9-1.2-2.9,0.6-2.2,2.3c0.8,2-0.1,4.8,2.7,5.9    c-2.2,2-3.3-0.8-5.2-0.9c0.5,2.7,1.4,5.1,3.8,7.4c-2.2,0.2-3.6,0.3-5.2,0.5c-0.2,2.9,2.8,3.9,3.7,6.2c-3.9,0.9-4.2,2-3.1,8.2    c2,11.4,7.1,21.2,15.6,29c6.5,6,10.6,13.8,16.6,20.1c3.3,3.4,9,6.8,13.7,5c6.8-2.7,9.2-5.6,9.6-12.1c0.6-9.1-3.9-15.9-9.4-22.1    c-6.4-7.2-13.8-13.6-19.5-21.4c-4.5-6-7.3-12.8-8.2-20.4c-0.5-4-0.8-7.9,0-11.9c0.7-3.9,2.5-7,5-9.7c2.3-2.5,4.2-5.1,5.4-8.4    c1.7-4.4,5.8-6.8,10-8.5s8.8-2.5,12.6-5.1c1.2-0.8,2.9-1.1,3.8,0c0.8,1-0.6,2.3-1.1,3.4c-0.4,0.8-2,1.1-1.3,2.3    c0.6,1.1,1.8,0.6,2.8,0.7c4.1,0.2,7.9-0.4,11.6-2.2c4.5-2.2,5-4,2.1-7.7c3.1,0,5,1.1,4.5,4.4c-1,7.4-3.3,14.3-9.8,18.7    c-2.5,1.7-5.8,1.6-8.8,3.7c3.8,0,7.2,0.8,9.8-2.1c0.7,4.4-3.1,7.9-8.5,8.4c-4.8,0.4-8.8-1.4-12.8-3.6c-1.3-0.7-2.6-1.3-4.2-1.8    c1.7,2.1,3.5,4.2,1,6.7c-0.5,0.5-0.9,1.3-1.9,0.6c-5.4-3.7-7.3-3.1-10.6,2.6c-3.7,6.4-1.5,12.6,1,18.5c2.3,5.3,10.1,14.5,13.8,17    c0.4-3.8,1-7.4,2.8-10.7c2.3-4.4,6.2-6.6,10.8-7c2.5-0.2,4.9-1,7.5-0.9c2.5,0.1,3-2.1,3.8-3.6c0.5-1,0.9-1.8,2.1-1.7    c1.7,0.1,0.9,1.5,1,2.4c0.6,8-0.9,10.6-9.1,12.5c1.4,1.2,2.8,2.2,4,3.5c1,1.1,0.9,2.7,0.2,3.9c-0.6,1-1.6,0.9-2.7,0.5    c-4.6-1.7-8,1-8.5,6.2c-0.5,5.5,2.7,9.6,5.4,13.6c6.4,9.7,6.8,19.8,2.9,30.3c-0.4,1.1-0.4,2.2-0.6,3.3c-0.3,1.4,1.1,4.2-2.4,2.9    c-0.4-0.1-0.6,0.1-0.6,2c0,1.8-0.4,3-2.3,1.8c-3.7-2.4-4.2-0.6-3.9,2.7c0.1,0.7-0.1,1.4-0.1,2.1c-1.2,0.5-1.6-0.4-2.1-0.8    c-3.4-2.8-3.9-2.7-4.7,1.4c-0.4,2.2-1.4,2.5-3,0.9s-3.3-3.2-5.1-5.1c-1.9,4.2-3.6,3.2-5.8,0.1c-1.3-1.9-3.7-3-4.9-5.4    c-0.6,5.3-0.5,10.5-1.1,15.6c-0.6,5.4,2.2,9.6,6.5,13c2.1,1.6,3.7,3.7,6,5.2c1.9,1.2,1.9,4,0.8,5.1c-2.9,2.8-4.5,7-9.1,8    c-1.8,0.4-3.1,2-4.5,3.4c-2.8,2.8-5.6,5.6-8.4,8.5c-2,2-3.6,4.2-3.1,7.3c0.1,0.8,0,1.6-0.5,2.5c-0.8,1.5-1.5,2.9-3.6,2.9    c-2.6-0.1-1.9-2.2-2.3-3.9c-1.7,1.9-2.4,3.1-0.9,5.8c2.1,3.9-1.5,6.9-3.9,9.4c-2.3,2.3-5.1,4.3-8.3,5.7c-1.1,0.5-2.3,1.1-3.1,0.4    c-0.9-0.8-0.7-2.2,0.1-3.2c-1.2-0.8-1.4,0.3-2,0.6c-1.8,0.9-4,1.2-5.3,0c-1.2-1.2,1.1-2.2,1.7-3.4c1.5-3.2,2.8-6.6,4.3-9.8    c1-2.2,2.1-4.4,4.5-5.5c1.7-0.7,3.3-2.1,3.8-3.6c1.3-4.4,5-6.3,8.4-8.6c4.3-3,7.5-7.1,10.2-11.5c2.8-4.6,1.5-7.4-3.8-8.5    c-2.4-0.5-4.4-1.6-6.4-2.9c-1.3,3.3-3.3,6.2-2.8,10c0.2,1.5-1.1,2.4-2.8,2.8c-0.1-1.5-0.2-2.9-0.3-4.6c-1.1,2.2-1.9,4.3-3.6,5.5    c-2.1,1.5-4,0.8-5.4-1.4c3.7,0,4.8-1.7,3.3-5.7c-1.8-4.9-4.2-9.7-3.4-15.1c0.7-4.3,0.1-8.3-0.7-12.5c-0.9-5.1,0.7-10.1,2.5-16.2    c-6.3,6.9-14.6,8-21.1,12.3c-1.6,1.1-3,2.4-4.1,3.9c-1.2,1.5-1.7,3.1-0.3,5.4c3,4.8,4.7,10.3,8.6,14.7c1.8,2.1-1,6.1-4.5,6.8    c-9.4,1.8-18.8,3.6-28.2,5.4c-1.3,0.2-1.9,0.8-2.5,1.9c-0.9,1.7-2.5,2.8-4.3,3.5c-1.3,0.5-2.6-0.2-3.5-0.9    c-0.9-0.7-0.8-2.1,0.1-3.1c0.8-0.8,1.4-1.6,0.7-2.7c-0.3-0.5-0.9-0.2-1.3,0.1c-1.1,0.7-3.1,1.1-2.6,2.6c1.6,4.5-1.8,5.6-4.6,6.7    c-4.6,1.8-9.6,2-14.5,1.8c-1.5-0.1-2.4-1.1-2.8-2.2c-0.8-2.2,2-2.2,3-4c-2.4,0-4.4,0.1-6.5,0c-1.1-0.1-2.3-0.4-2.7-1.6    c-0.5-1.7,1.3-1.9,2.1-2.1c4.8-1.4,8.7-5,13.3-6.4c5.7-1.7,11.3-3.8,17.1-5.2c4.6-1.1,9.3,1.1,14,0.7c2.3-0.2,4.7-0.3,6.6-2    c1.9-1.6,2.3-3.4,0.2-5.1c-6.5-4.9-11-11.8-17.8-16.6c-0.6,3.4-2.9,4.2-6,3.4c0.2-1.2,2-1.8,1-3.5c-4.4,4.3-9.6,5.4-15.4,3.6    c-2.3-0.7-3.9-2.2-4.3-4.7c-0.2-1,0.4-1.9,1.3-2.1c1.2-0.3,0.9,0.8,1,1.5c0.2,2.6,1.9,1.5,3,1.1c3.3-1.2,5.9-3.6,7.6-6.5    c1.3-2.1,2.8-3.8,4.5-5.6c1.9-2,4-3.7,4.9-6.5c1.1-3.1,3.9-5,6.8-5.5c9.4-1.7,18.8-5.1,28.5-3.9c4.8,0.6,9.1,0.8,13.3-1.7    c3.7-2.2,3.5-5,0.3-7.9c-7.1-6.2-16.4-8.1-24.4-12.5c-5.3-2.9-9.9-6.4-12.6-11.9c-1-2.1-2.6-1.9-4-1.2c-2.6,1.3-5.4,1.6-8.1,1.2    c-6.8-1-13,1.1-19,3.6c-6.2,2.6-10.1,8.1-14.3,13.3c-0.9,1.2-0.4,2.5-0.5,3.7c-0.1,2.3-1.5,3.7-3.4,4c-2.2,0.4-2-1.7-2.1-3.2    c0-0.5,0.1-1.2-0.6-1.2c-0.4,0-0.9,0.3-1.3,0.6c-0.7,0.6-0.6,1.4-0.2,2.1c2.3,5.1,1.9,6.4-3,9.6c-7.9,5.2-12.2,6.9-16.9,6.6    c0-1.3,1.8-1.5,1.7-2.8c-0.1,0-0.3-0.1-0.4-0.1c-1.8-0.4-4.4,1-5.2-1.2c-0.8-2,2.1-2.1,3.1-3.2c2.6-3,4.3-6.7,7.1-9.6    c5.7-5.9,10.6-12.5,18.3-16.2c3.8-1.8,6.2-5.3,8.9-8.6c3-3.7,6.3-7.5,11.8-7.4c5.7,0.1,10.7-1.3,13.8-6.6c0.2-0.4,0.9-0.4,1.3-0.7    c0.3-0.2,0.5-0.6,0.8-0.9c-1,0-2.3-1.6-2.6-0.9c-1.5,4-6.1,2.6-8.5,5.1c-2.4,2.4-13.3,2.5-15.9,0.3c-1.9-1.6-3.8-3.4-2.5-6.4    c1.6-3.7,4.4-4.4,7.8-3.5c0.4,1-0.4,1.3-0.6,1.8c-0.3,0.6-1.9,0.8-0.9,1.9c0.6,0.8,1.4,1.8,2.5,1.3c3.2-1.3,6.4-2.8,8.1-6.1    c0.4-0.7,0.4-1.4-0.2-1.6c-4.6-1.3-3-5.6-3.9-8.5c-1.5-4.4-2.1-9.2-3.3-13.7c-0.8-3.3-2.8-6-4.6-8.8c-1.3-2.1-2.6-4.1-3.1-6.6    c-1.3-6.7-4.2-9-11-9.1c-1.3,0-2.8,0.4-3.4-1.1c-0.6-1.4-0.1-2.7,1.2-3.7c0.9-0.7,1.4-1.4,0.3-2.5c-1-0.9-1.8-1.9-3.3-0.7    c-2.9,2.4-5.6,1-7.1-1.3c-3.8-5.7-7-11.7-6.6-18.9c2.1-0.7,2.2,1.8,3.7,1.7c0.8-1.7-1.7-3.2-0.2-5c2.8,0.1,3.1,3.6,5.2,4.2    c8.5,2.6,12.5,9.9,17,16.5c2.5,3.6,5.7,6.4,9.4,8.5c5.9,3.3,7.7,9.2,9.5,15c0.8,2.8,1.2,5.5,3.2,7.7c3.1,3.6,5.1,8,8.4,11.4    c2.4,2.4,4.3,3.8,7.5,0.5c2.6-2.7,6.5-3.6,10.2-3.9c5.5-0.4,11-0.6,16.3-2.4c4.1-1.4,8.2-3,10-7.6c0.6-1.5,1-2.5-0.9-3.2    c-3.1-1.1-5.2-3.4-7.5-5.6c-3.5-3.4-12.5-4.7-17-2.9c-1.4,0.6-1.7,1.8-2.2,2.9C123.9,118,121.6,119.6,117.8,118.9z"
	pathData1 := "M178.6,132.2c1.2,2.9,0.6,6.1,1.2,9.5c-1.5-1-1.9-2.4-2-4c-1.1,0.8-0.3,1.8-0.9,2.6c-1.1-0.1,0-1.4-1.4-1.7    c0.5,3.4,1.1,6.7,1.6,10.1c-2.2-1.8-2.2-1.8-2.3-9c-1.4,1.5-1.4,1.5-0.8,4.3c-1.4,0.5-1.1-1.2-2.1-1.6c-0.4,2.6,1.8,4.6,1.4,7    c-0.4,2.6,2.5,4.3,2.2,7c-2.1-0.1-1.9-2.5-3.9-3.5c0.2,4,0.6,7.5,3.7,11.6c-5.1-3.3-6.5-0.3-7.9,2.9c1.7,0.4,1.9-1.9,3.5-1.9    c-0.3,1.9-2.3,2.6-3,4.2c1.9,0.5,2.1-1.8,3.9-1.9c-0.9,2.2-2.3,3.9-3.8,5.6c1.4,1.2,1.7-1,2.7-0.9c0.9,1.5-1.8,1.7-1,3.5    c0.9-0.8,1.7-1.4,3.3-2.9c-1.1,3.1-3.2,4-4.3,5.8c1.8,0.6,2.1-1.5,3.5-1.8c0.6,1.9,1.2,3.7,1.8,5.6c-1.8,0-2.3-3-3.1-2.1    c-1.3,1.4,0.7,2.7,1.8,3.8c2,2.2,2.5,5.4,5.7,6.8c1.1,0.5,0.9,3.4,2.4,4.7c0.1,0.1-0.3,0.8-0.2,0.8c0.6,0.4,2.5-1.1,1.8,1.1    c-0.3,0.9-1.6-0.8-2.5,0c0.6,1.6,2.3,1.1,3.5,1.8c1.7,0.9,3.2,1.7,2.3,4.1c-0.3,0.7-0.4,1.7,1.2,1c1.4-0.6,0.5,0.9,0.5,1.2    c0,2.6-0.1,2.6,2.5,1.8c0.1,0,0.2-0.1,0.4-0.1c1.5,0.1,1.6,3.2,3.9,1.9c-0.1,1.2,1.4,2.1,0.5,3.5c-0.8,1.3,0.6,0.6,0.9,0.7    c1.9,0.8,2.5,2.7,3.9,4c-3.2,1.1,0.7,2,0.2,3.2c-0.1,0.2,0.6,1.2,1.2,0.9c0.3-0.2,0.3-1.1,0.8-1.6c0.1,1.1,0.2,2.1,0.3,3.2    c1.2-0.1,0.2-1.3,1.1-1.8c0.6,1.8,2.1,3,3.3,4.5s1.8,3.7,1.8,6.3c1.2-2.4,0.2-4.3-0.1-6.3c2.4,0.8,0.9,3,1.8,4.2    c2.3-0.4-0.3-1.9,0.8-2.6c2.3,2.4,4.6,4.2,7.9,1.5c0.5-0.4,1.5-0.1,2.2-0.3c11.1-2.9,13.9-12.5,10.8-22    c-2.2-6.6-5.9-12.6-10.7-17.7c-6.8-7.1-14.2-13.9-19.6-22.2c-3.6-5.6-5.9-11.7-6.5-18.7c-0.7-8.2,1.6-14.6,7.5-20    c1-1,1.7-2.2,2-3.5c1.8-6.9,7.1-10.3,13.2-12.3c3.5-1.2,7-2.3,10.3-4.4c-0.9,3.5-2,4.1-6.9,5.8c-3,1-6.1,1.9-8.8,3.8    c0.8,0.9,1.4,0.6,2.2,0.2c4.6-2.2,9.4-3.2,14.5-2.8c4.2,0.3,8.4,0,12.7-1.2c2.2-0.7,3.5-1.6,4.9-3.6c-3.4,12.6-8.7,14.5-17.4,13.2    c-2.1-0.3-4.3,0.1-6.3-1.1c2,1,4.2,0.9,6.4,1c1.3,1.1,6.1,1.6,8.8,0.7c0.6-0.2,1.2-0.4,1.8-0.3c0.1,0,0.2,0.1,0.3,0.1    c-2.7,2.6-6.4,3.4-9.7,4.1c-3.1,0.8-6.3-1.8-9-3.6c-3.2-2.1-5.6-0.7-8.4,0.6c5.8-0.4,10.8,1.3,14.8,5.6c2.3,2.4,5.4,1.3,8.4,1.7    c-2.4,3-5.3,0.6-7.9,0.8c0.7,1.5,2.5,1,3.5,2.1c-3.8,0.7-7.2-0.7-10.1-2.3c-4.5-2.6-8.9-3.8-13.9-1.7c1.4,2.2,6.3,0.1,6.4,4.8    c-3.1-2-6.4-2.2-9.2-0.5c-3.4,2.1-5.8,5.4-6.2,9.7c-0.9,10.7,3.4,19.5,10.4,27.1c2.8,3,5.3,6.5,9.4,8.1c0.7,0.3,1.4,0.6,2.5,0.1    c-3.3-4.7-1-9.1,1.1-13.4c2.1-4.4,5.8-6.2,10.6-6.3c4.1-0.1,8.6,0.3,11.1-4.5c1,4-1.2,4.7-4.3,5.4c-2.8,0.6-6.2-0.7-8.3,2.4    c-1.1,0.1-2.2,0.3-3.1,0.9c0.9-0.6,2-0.7,3-1c3.1-0.9,6.4-0.4,9.5-1.6c0.5-0.2,1.5,0,1.7,0.3c0.5,1.1-0.8,0.9-1.2,1.4    c-2.8,2.8-6.7,2.4-10,3.6c-2.7,1-4.8,2.4-3.7,6.1c2-2.7,4.2-5.3,7.8-4.3c2.1,0.6,4.2,1.9,4.3,4.9c-3.2-1.1-6.3-2.8-9.2,0.4    c-3.3,3.6-3.9,7.7-2.6,12.2c1.5,5.3,4.7,9.8,7.2,14.6c4.3,8,4.1,15.9,0,23.9c-0.3,0.6-1.2,1.1-0.6,2c0.7,1,2.3,1.9,0.8,3.1    c-1,0.8-2.4,0-4.2-0.9c1,1.7,2.2,3,1.2,4.9c-2.4-0.8-4.9-1.2-7.1-2.9c-0.8,2.2,2.2,3.6,0.9,5.8c-1.4-1.3-3.3-1.7-4.9-2.5    c-1.6-0.8-2.3,1-3.5,0.5c-0.1,1.1,0.1,1.8,1.3,2.4c0.8,0.4,0.6,1.7,0,2.9c-1.6-1.8-2.4-4-4.7-4.9c-0.6-0.2-1.1-1.6-0.9-2.3    c0.4-1.1,1.1-0.3,1.8,0.1c0.2,0.1,1.1,0.4,0.8-0.5c-0.2-0.5-0.6-1-0.9-1.4c-0.4,1.7-2.2-1-2.6,0.8c-0.2,0.7,0.5,2.3-0.9,1.9    c-1.5-0.4,0-1.8-0.3-3c-1.8,0.7-1,2.1-1,3.1c0,0.8,1.5,1.9,0,2.5c-1.4,0.6-1-1.1-1.6-1.7c-2.5-2.1-5.2-4-4.7-7.9    c-1.3,0.8-1.4-0.6-2.1-0.9c-0.1-0.1,0.1,1.2-0.5,1.5c-1.4-1,1-2.8-1-4.1c-0.8-0.5-1-0.4-1.6-0.3c-0.8,0.2-1.2,0.7-0.7,1.3    c3,3.8-0.1,8.2,1.6,12.2c-1.6,0-0.6-1.4-1.4-1.6c-0.7,0.5-0.6,1.7-0.3,2c1.5,1.6-0.5,2.4-1,2.6c-2,0.6-1.7,2.7-1.3,3.2    c2.4,2.8,2.4,5.9,2.1,9.3c-1-0.1-0.8-1.8-2.1-1.6c-0.3,2.6,1.3,5,0.9,7.7c0.4-2.5,0.8-4.8,3.6-6.4c-0.2,3.3,0.7,6.5-1,10    c3-1.7,0.9-4.4,2.5-6.2c0,2,0,3.8,0,5.8c1.5-0.6,0.5-1.8,1.2-2.6c1.5,3.1,4.1,5.7,3.7,9.7c1.5,0.2,2.4-1,2.8-2    c0.4-1.1-0.9-2-1.7-2.3c-1.7-0.6-1.4-2-1.7-3.3c0.5,0.5,1.1,1,1.6,1.5c5.3,5.2,4.7,9.8-2.7,12.4c-5.6,2-8.6,6.3-12.1,10.2    c-0.9,1-1.7,2.1-2.6,3.2c-2.3,2.9-4.5,5.7-3.5,9.8c0.1,0.5,0.1,2.2-1,2.8c-1.1-0.1-0.8-0.9-0.9-1.5c-0.4-2.8-1.7-3-3.8-1.3    c-2.6,2.1-3.3,6.1-7.2,6.9c-0.7,0.1,0.3,0.5,0.6,0.6c0.7,0.2-0.8,2.1,0.9,1.3c0.7-0.4,1.3-0.7,2.6-0.4c-2,1-1.4,3.1-3.2,4    c1.7,0.9,1.8-1,3.5-0.7c-3.4,2.9-5.4,6.7-9.6,9.3c1.1-5.1,2-9.5,3.1-14.6c-5.2,2.3-2.7,9.6-8.6,10.7c3.8-4.9,2.8-12.6,9.6-16    c0.7-0.3,1.4-0.9,1.7-1.5c2.3-5.8,7.6-8.7,11.9-12.7c4-3.7,7.4-7.8,9.2-13.1c0.8-2.4,2.6-0.1,3.3-1.3c-1.1-0.7-1.1-1-0.1-2.1    c-5.7-2.1-11.8-3.9-16.9-8c-1.3-1-1.2-2.9-1.8-4.4c1.4-0.7,0.2-1.3,0-2c-0.2-0.5-0.6,0.3-0.2,0c0.1,0,0.3,0.3,0.4,0.5    c0.3,0.6,0.5,1.1-0.1,1.6c-0.4,0.3-1.8-1.4-1.3,0.3c0.4,1.4,0.5,2.7,0.6,4.4c-1.2-1.5-2.2-2.5-3.2-3.8c1.3,1.4-0.8,3.8,2.2,4.8    c1.8,0.6,1.4,2.5-0.6,3.1c-1.2,0.4-0.7,1.4-0.6,1.6c0.7,0.8,0.7,1.6,0.6,2.5c0,0.5-0.1,1-0.1,1.7c-2.1-1.4-2.1-4-4.2-6    c0.2,3.7,1.6,6.9-1.1,10.1c0.7-4.7-1.4-8.2-2.8-11.9c-1.4-3.8-1.8-7.8-0.1-11.3c1.2-2.5,0.2-4.6-0.3-6.5c-1.4-5,0.1-9.7,0.7-14.5    c0.7-6,2.4-12,0-17.9c-3.8-9.3-11-15.2-20.2-18.7c-7.9-3-15.4-6.9-22.1-12.8c2.9-0.3,3.6-2.7,4.7-4.3c3.1-4.1,8.1-4.7,11.8-7.6    c0.7-0.6,2.5-1,0.7-2.4c-0.5-0.8,0.2-1.7-0.2-2.6c-0.1-0.2-0.2-0.4-0.3-0.6c-0.6,0.4-0.1,1.1-0.3,1.6c-0.4,1.1-0.8,2.2-2.3,1.7    c-1.7-2.1,2.4-3.8,0.1-6.3c-0.6-0.6,1,3.8-1.1,1.1c-0.7-0.9-1.2,0.2-1.2,1.1c0,2.1-1.4,3.4-2.6,4.8c-1.4,1.6-3,2.7-3.8,5    c-1.1,3-6.7,3.9-10.3,1.4c-1.8-1.2-2.9-0.5-4-0.2c-3.2,1.1-6.2,1-9.6,1c-7.9-0.2-16.2,0.9-22.5,5.7c-5.2,4-11.2,8.3-11.1,16.4    c0,0.9-0.6,2.2-1.9,2.8c-0.7-1.6-0.5-3.9-3.4-3.7c-3.5,0.2-5.1,2-4,5.4c0.7,2.2,0.2,3-2.1,3.8c0.9,0.1,1.6,0.1,2.2,0.1    c0.5,0,0.9,0.1,1.4,0.6c-1.3,0.1-2.7,0.2-4,0.3c2.5,1.7,0.5,3-0.6,3.7c-2.9,1.9-6.3,2.5-9.3,4.4c1-1.9,1.7-3.9,2.9-5.6    c1.2-1.8,2.2-3.5,1.9-6c-3.8,2.8-4.1,8.6-10.2,9.7c5.5-5.2,8-11.7,13.8-15.8c3.5-2.5,5.6-6.8,9.4-9.1c2.3-1.4,4.6-3,6.9-4.4    c2.8-1.8,4.8-4.2,6.6-7c2.1-3.3,5.2-6.7,9.3-6.7c6.4,0,11.7-1.3,16.4-6.1c3.6-3.7,9.6-2.6,14.1-5.6c-4.6-0.3-8.4,2-12.9,3    c1.2-2,2.6-3.6,3.3-5.4c0.6-1.6,3.7-2.8,0.8-5.4c-0.3,3.8-2.2,6.1-5.5,8c0.6-2.3,3.6-2.8,2.8-5.6c-3.2,2.9-6.2,5.7-9.4,8.6    c-1.3-3.1,2.7-4.8,1.7-7.8c-1.7,1.4-2.2,3.4-3,5.1c-0.9,1.9-2.3,3.2-4.2,4.8c-0.2-1.1-0.3-2-0.4-3c-2.9,3.8-9.3,5.3-12,2.7    c-1.9-1.8-4.1-3.9-0.6-6.4c1,4.5,3,5.3,7.2,3.7c4.5-1.6,7-5,9.2-8.8c2-3.4,4.9-5.6,8.6-6.7c2.8-0.8,4.6-3.1,6.6-4.8    c3.7-3.1,7.8-4.1,12.3-4.5c10.5-0.7,20.7-2.4,27.1-12.4c1.3-2,3-2,6.4-0.8c-1.1,0.1-1.8,0.1-2.5,0.1c-0.6,0-1.2,0.1-1.8,0.3    c2.9,0.9,5.5,0.2,8-1.2c-1-1.7-1.8,0.8-2.9-0.3c1-1.2,3.7-1.9,2.4-3.9c-0.7-1-3-0.8-4.5-1c-2.8-0.3,1.1,3.2-1.8,3.2    c-0.6-0.9-1.1-2.3-2.8-2.5c-1.4,1,0.9,1.8-0.1,2.9c-0.4-0.4-0.7-0.9-1.2-1c-1.2-0.4-0.8-5.1-3.5-1.3c-0.5,0.6-1.5-0.7-2-1.3    c-3.6-4.7-8.9-6.3-14.2-7.6c-3.6-0.9-7.1,0.9-10.7,1.5c-0.3,0.1-0.8,0.6-0.8,0.7c1.5,3.7-3.7,3-3.6,5.8c-1.1-1.9-0.6-3.9-1.1-5.8    c-0.6-2.3,1.5-3.7,3.9-4c2.3-0.3,4.6-0.8,5.4-3.5c3,2.3,11.6,2.8,14.4,0.7c1.6-1.2,2.2-2.7,0.6-4.5c-2.5-2.9-6-4.5-9.3-6.2    c-1.9-1-2,1.6-3.4,1.7c-0.9-4.5-2.2-4.7-6.6-1.1c-1.4-2.5-2.4-6.2-6.1-2.1c-0.7,0.8-1.1-0.2-1.4-0.6c-1.4-1.9-4.7-2.8-3.9-6    c0.3-1.1,0.6-2.1,0.7-3.2c0.1-4.7,0.7-4.8,5.2-4.1c7.2,1.2,14,4.9,21.7,3.6c1-0.2,2-0.1,3-0.3c2.1-0.4,4.3-2.8,6.2,0.3    c0.4,3.1-2.2,4.2-4.3,5.9c3.7,0.4,5.5,5.6,10,3c0.3-0.1,0.7-0.2,1.1-0.2c2.4,0.8,2,2.3,1,3.9c-0.7,1-1.7,1.7-2.4,2.5    c1.8-1.3,2.8-4,6.4-3.6c-3.3-1.7-5.5-3.3-5.8-6.8c-0.2-2.9-2.7-4.3-5-5.6c-0.1-3.1,0.6-6.6-4.2-7.2c2.5-1.4,6.3-0.7,7.2,1.1    c3,5.9,3,5.9,8.3,1.9c0.9,1.5-0.7,2.2-1.2,3.3c2.7,0.9,3.4-1.8,4.9-2.8c5.9-4.3,12.9-5.5,19.8-6.9c12.7-2.4,25.6-3.4,38.1-7    c2.8-0.8,5.5-2,9.2-3.4c-4.7,4.9-10.4,6.2-15.7,9.1c3.6,1,5.3,3.3,6.5,6.4c-6.7-6.4-14.9-4-23.4-4.5c3.1,3.2,5.5,5.7,8,8.2    c-6.4-4.7-13.6-5.9-21.7-4.8c1.1,2.3,3,3.4,4.4,5c0.5,0.6,0.6,1.1,0.1,1.7c-0.4,0.5-0.8,0.8-1.5,0.1c-4.3-4.4-9.7-3-15.2-3.1    c0.7,2.2,1.1,4.2,3.9,3.8l-0.1-0.1c0,0.8,0.3,2.1-1.2,1.4c-0.9-0.4-1.6-0.5-2.5-0.5c-0.6,0.5-1.1,0.4-1.7,0    c-1.9-1.8-3.7,1.2-5.7,0c1.4,1.3,2.8,2.3,5.6,2.5c-2.5,0.7-4,1.1-5.6,1.6c2.5-0.8,5.2,1.3,7.8-0.8c1-0.8,2.9-0.8,4.4-0.9    c9.4-0.7,18.8-0.7,28.3-0.8c14.4-0.2,27.9-4.1,40.3-11.6c1.8-1.1,3.1-2.9,5.3-4.9c-1.3,4.5-4.1,6.9-7.1,8.8    c-2.6,1.6-5.3,3.5-8.6,4.1c-1.5,0.3-3.2,1.2-3.7,2.9c-0.6,2.1,1.6,2.5,2.7,3.1c2.2,1.2,3.1,3.5,4.7,5.5    c-7.3-4.5-15.3-5.3-24.2-4.4c3.4,2.9,6.4,5.2,7.9,8.9c0.4,1,0.7,2.4,0,2.8c-1.1,0.7-1.4-0.6-2.1-1.4c-5.6-7.2-13.6-8-22.3-8.5    c1.3,2.7,4.3,3,5.3,5.2c0.8,2,2.6,3.8,1.4,6.2c-2-4.9-10.6-11.5-18.5-10.4c0,2.5,3.2,3.8,2.5,7c-2.7-4-6.4-5.2-10.8-5.5    c-4.2-0.3-8,0.9-12.8,2.7c3.5,0.8,5.6-0.5,7.9-0.9c2.6-0.4,5.2-0.6,7.8,0.3c2.3,0.9,4.9,1.1,6.6,4.1c-3.7-0.8-6.9-1.7-10.2-2.1    c-3-0.4-6.4-1.5-9,1.4c7-1.4,13.5,1.4,20.2,2.4c3.2,0.5,4.8,3.7,6.8,7c-3.8-2.1-7.2-3.9-11.3-6.2c1.1,2.6,3.3,3.3,3.4,5.5    c-3.2-0.9-7.1,1-9.2-3.2c-0.5,0.8-0.7,1-0.7,1.2c-0.1,1.2,2.3,1.7,1.2,2.8c-1.2,1.3-2.1-0.5-3.2-1c-1.9-0.8-3.2-2.2-4.1-3.8    c-0.7-1.1-1.2-1-2.3-0.8c3.3,5.6,7.1,8.7,14.2,7c3.2-0.8,7.1,0.7,11-0.6c-0.6,1.5-1.1,2-2,2c-5.3-0.2-10.5,1.3-15.7,0    c-3.8-0.9-6.1-3.9-8.4-6.7c-0.8-0.9-1.1-2-2.5-1.7c-0.4-0.4-0.6-1-1.1-1.3c0.2,0.6,0.8,0.8,1.3,1.2c-0.2,1.1-1.3,1.4-2.2,2    c2.6,0.3,3.7,2.4,5.3,3.9c2.2,2,4.2,4,7.6,4.3c3.4,0.3,3,3.9,3.7,6.1c0.4,1.1,2.1,4,0,4.3c-2.7,0.3-0.9-2.6-1.3-4.1    c-0.6-2-0.6-4.2-3-5.6c-0.1,1.2-0.2,2.3-0.3,3.3c-0.1,0-0.3,0-0.4,0c-0.2-1.1-0.3-2.1-0.5-3.2c-0.3,2.5-2.6,4.9-0.2,7.5    c0.7,0.8,1.2,2.4,0.9,3.4c-0.4,1.6,0.7,2.1,1.3,1.7c1.9-1,1.4,0.8,1.2,1c-1.1,1.1-0.3,2,0,3c0.1,0.4,1.3,1,0.2,1.3    c-0.4,0.1-1.1-0.5-1.6-0.8C178,130.2,178.8,131.4,178.6,132.2z"
	pathData2 := "M122.7,245.4c-2.8-0.5-4.2-2.9-5.7-4.8c-1.2-1.5,0.2-3.5,1.5-4.6c3-2.5,3.1-2.5,1.7-6.2    c-1.6,0.8-0.1,2.6-1.1,3.5c-1.5-0.2-1-2.1-2.5-2.5c-0.4,2.1-0.8,4.1-1.4,6.8c-0.3-1.3-0.5-2.1-0.8-3.2c-1,3-1.7,5.9-4.2,8.4    c-0.1-2.9-0.3-5.5,1.2-8.9c-3.4,2.9-4.5,6.7-7.8,8.5c-0.9-2.1,2.4-2.8,0.7-4.8c-2.8,3.3-4.9,8.5-11,4.7c5.4-1.8,9.3-5.3,11.9-10.2    c0.8-1.4,2.1-2.7,3.5-3.5c2.5-1.5,3.5-3.7,4.6-6.3c1.5-3.4,5.2-4.3,8.7-4.8c7.3-1,14.1-4.2,21.6-3.7c3,0.2,6,0.5,9,0.9    c1.8,0.2,3.3,1.1,4.2,2.8l-0.1-0.1c-0.2,1.3-0.5,2.3,1.5,1.2c3.4-1.9,6.1-0.8,8.3,2.6c0.7,1.1,1,1.2-0.1,1.9c-1,0.5-2,1-2.9,1.6    c-4.1,2.4-7.4,5.8-10.7,9.1c-1.4-1,1-1.7-0.1-2.8c-1.1,1.6-2.3,3.1-3.4,4.7c-0.2-0.1-0.3-0.2-0.5-0.3c0.4-1.1,0.9-2.2,1.3-3.4    c-2.5,2.8-5.2,5.2-7.8,7.7c-0.3,0.3-0.5,1.6-1.5,0.6c-1.9-2-3.1-0.8-3.8,1.2c-0.9,2.4-0.5,4.4,1.5,6.5c1.3,1.4,3.9,2.5,2.9,5.5    c-0.1,0.3,1.2,1,2.1,1.7c-1.2,0.3-1.8,0.5-2.3,0.6c2.9,2.3,4.4,5.8,8.4,8.5c-3.5,0-4,2.6-6.6,4.2c1.2-2.7,3.3-4.4,1.8-7.3    c-1.9,3.1-3.5,6-7.3,7.8c1.6-2.3,2.8-3.9,3.9-5.5c-5.8,4-11.5,8.1-18.8,8.9c0.1-1.9,2.9-1.3,2.7-3.7c-5,3.4-10.7,5.4-14.5,10.3    c-0.4-2.3,2.8-3.3,0.5-5.8c-2.8-2.9-5.6-1.6-6.6,0.6c-2,4.2-5,1.9-7.6,2c-0.5,0-1-0.6-1.8,0.4c2.4,1.6,5.2,0.1,7.7,1.1    c-2.1,1-4.8-0.5-6.5,1.8c1.9,0.1,3.7,0.2,5.5,0.3c0,0.2,0,0.4,0,0.6c-2.3,0-4.7,0-7,0c0,0.2,0,0.4,0,0.6c1.7,0.1,3.5,0.2,6.4,0.3    c-5.3,2.3-9.9,3.4-14.6,2.8c-2.3-0.3-0.6-1.9,0.2-2.6c2.9-2.3,6.4-4.1,8-8.2c-4.3,2.7-8.3,6.4-14.1,4.4c5.3-2.3,10.4-4.7,15.2-7.6    c1.2-0.7,2.5-0.5,3.7-0.4c2,0.2,4.3-0.1,5.4-1.2c3.8-3.7,7.9-2.4,12.4-2.2c4.6,0.2,9.7,1.9,14.1-1.3c3.8-2.7,4.3-5.6,0.8-8.6    c-2-1.7-2.7-4.9-5.9-5.4c-1.5-0.2-0.9-1.8-0.8-2.3c0.8-2.1,2.2,0.8,3.9-0.8c-2.5-1.3-5.8,0.5-7.3-2.1c0.1-0.6,0.9-0.1,1-0.7    c0,0-0.1,0-0.3,0C122.7,244.7,122.8,245.2,122.7,245.4z"
	pathData3 := "M65.1,93.3c-2.4,1.2-4.2,1.1-5.1-2c-1.1-3.6-4.8-6.1-4.5-10.5c3.8,2.1,6.8,5.5,12.4,6c-4-2.9-9.2-3.4-9.5-9.5    c2.4,2.6,5,3.8,7.7,5c4.5,2,5.8,6.8,8.6,10.4c3.2,4.2,6.4,8.4,11,11c5.5,3,8.6,7.8,9.6,13.7c1.1,6.8,5.8,11.3,9.2,16.7    c1.9,3,4.8,5.2,7.3,8.3c-4.6,1-7.9,3.6-10.7,6.9c-0.8,1-1.9,0.8-2.8,0.6c-1.1-0.3-1.2-1.3-1.4-2.4c-1.6-10.1-3.2-20.2-9.2-29.1    c-2.3-3.4-1.8-7.8-4.5-11.2c-3.2-4.1-4.7-4.9-9.7-3.6c-1.1,0.3-2.2,0.2-3.3,0.2c-0.5-1.6,0.6-2.1,1.3-2.7c1.1-1,1.9-2.2,1.4-3.7    s-2.3-2.4-3.3-2C67.8,96,66.7,93,64.9,95C65,95.1,64.8,93.9,65.1,93.3c-0.2-1.5-1.1-2.8-0.7-4.4C64.1,90.5,64.9,91.8,65.1,93.3z"
	pathData4 := "M141.5,103.7c-9.1-5-17.8-1.5-26.4,1.1c-3.7,1.1-6.7-0.3-9.6-1.6c-3.3-1.5-3.1-3.8,0-6.3    c1.8,4.6,5.4,5.9,11.5,4.2c6.8-1.8,13.5-3.1,20.5-0.6C139.4,101.1,140.5,102.1,141.5,103.7z"
	pathData5 := "M156.3,91.5c-1.5,1.7-3.2,2-5.2,0.7c-2.1-1.4-3.7-3.8-6.8-3.3c-0.2,0-0.7-0.6-0.8-1c0-0.2,0.5-0.8,0.7-0.8    c3.4,0.1,4.5-2.5,5.9-4.8c0.3-0.3,0.6-0.6,0.9-0.9c2.8-1.3,3.6,1.5,5.1,2.7c0.8,0.6,1.2,2.1,1.1,3.4c-0.2,3.4,2.1,4.4,4.9,4.8    c0.4,0.1,1.2-0.3,0.9,0.4c-0.3,0.7,0.1,1.7,0,1.6c-4.1-1.2-5.7,3.2-8.8,4c-0.3,0.1-0.7,0-1.6-0.1c2.2-0.7,3.6-2,4.6-3.5    c0.8-1.2,2.4-3.1-1-3.1c-1.2-1.6-1.2-3.6-1.8-5.4c-0.9-2.7-2.1-2.6-3.5-0.3c-0.8,1-1,2,0.1,3C152.1,90.9,154.8,89.9,156.3,91.5z"
	pathData6 := "M112.2,78.3c0.2,2,1.2,3.7-2.4,3.9c-3.7,0.2-6.7,2.3-8.6,5.9C100.6,82.7,101.7,81.7,112.2,78.3z"
	pathData7 := "M178.7,82.1c-0.4-1-1.9-2.2,0.8-2.4c3.2-0.2,6.1,0.6,9.6,2.7c-4.1,0.9-7.4,2.8-10.6-0.5    C178.6,82,178.7,82.1,178.7,82.1z"
	pathData8 := "M143,103.2c-2.4-2.9-5-4.9-8.7-5.1c-0.3,0-1.7-0.2-1.4-1.2c0.2-0.6,1-1.4,1.6-1.5    C136.7,95.1,142.2,100,143,103.2z"
	pathData9 := "M137.8,104.8c-4.6,0.8-5,0.7-7.5-1.9C132.9,101.9,136.4,102.7,137.8,104.8z"
	pathData10 := "M141.6,167.6c1.1-0.4,1.6-1.3,1.7-2.4c0.4-0.4,0.7-1.3,1.6-0.6l-0.1-0.1c0.9,1.5-0.4,2.6-0.8,3.9    c-1.5-0.4-2,1.3-3.2,1.7c-0.8,0.3-1,2-2.7,1.2C138.8,169.5,140.5,168.8,141.6,167.6z"
	pathData11 := "M168.4,218.7c-0.5-1.1-2.9-1.2-1.9-3c0.2-0.3,1-0.5,1.4-0.3C169.4,216.2,168.4,217.5,168.4,218.7z"
	pathData12 := "M178.6,132.2c-0.5-1.1-1-2.2-1.9-4c1.3,0.8,2.3,0.7,2.7,1.6C179.4,130.7,179.6,131.6,178.6,132.2z"
	pathData13 := "M173.2,82.9c0.6,0,1.1,0,1.7,0c-0.2,1,1.5,0.4,1.2,1.3c-0.1,0.3-0.7,0.7-0.8,0.7    C174.4,84.5,173,84.5,173.2,82.9z"
	pathData14 := "M143.9,168.3c0.3-1.3,0.5-2.6,0.8-3.9c2.6,0.9,0.1,2,0.1,3C144.5,167.7,144.2,168,143.9,168.3z"
	pathData15 := "M157,217.4c1.1-0.7,2.1-1.7,3.8-1.2C159.5,217.3,158.4,218,157,217.4C156.9,217.3,157,217.4,157,217.4z"
	pathData16 := "M166.3,213c-0.3-0.2-0.6-0.4-0.7-0.7c-0.1-0.7,0.4-1.1,0.9-1.4c0.6-0.3,0.5,0.3,0.5,0.6    C167.2,212.1,167.1,212.7,166.3,213z"
	pathData17 := "M170.3,112.2c-2.3,0.7-3.6,1.1-5,0.6c0.6-1.2,2,0,2.7-1c0.8-1.1,2,0.4,2.9-0.1c0.3-1.3-1.4-0.7-1.6-1.9    c-0.1-1-2.7,1.3-1.9-1c0.2-0.5,2.5-1.1,4-0.4c2,1,2,3.8,4.2,4.6c0.1,0,0,0.7,0,1.1c-2.4-0.5-3.9,0.6-4.9,2.6    c-1.8-2.5-3.5,0.5-5.3,0.2c-1-0.2-2-0.6-1.9-1.5c0.2-1.8,1.1-0.5,1.8-0.2C167.4,115.8,168.4,114.4,170.3,112.2z"
	pathData18 := "M166.6,144.1c2.4-3.8,0.6-6.9,0.5-10.3c2.3,1.6,1.1,4.1,2,6.7c1.2-2.3-0.3-5,2.2-6.4    c-0.1,1.5-0.3,2.6-0.3,3.6c0,0.6,0,1.1,0.9,1.4c0.8,0.3,0.5-2.3,1.4-0.5c0.7,1.4-1.8,1.1-1.6,2.5    C169.2,140.3,168.7,142.7,166.6,144.1z"
	pathData19 := "M165,154.7c-0.1-0.9-0.1-1.8-0.3-2.6c-0.2-0.7-0.8-1.5,0.3-1.8c1-0.2,0.4,0.8,0.6,1.3    c0.7,1.1,0.9-0.1,1.3-0.3c-0.7-1.8,1.3-3.3,0.5-5.1c0.6,0.8,2.5,0.8,1.3,2.6c-0.6,0.9-0.3,1,1.9,0.7c-2.1,1.9-3.9,3.4-5.7,5    L165,154.7z"
	pathData20 := "M217.4,120.5c5.8-2.4,11.9-0.8,17.8-1.1C229.3,120.4,223.3,119,217.4,120.5z"
	pathData21 := "M147.1,106.8c0.5-0.4,1-0.7,1.5-1.1c1-1,1.6-0.2,2.1,0.5c0.7,0.9-0.1,1.5-0.7,1.9c-2,1.3-4.2,1.2-6.8,0.8    C144.6,108,146.4,108.2,147.1,106.8C147.2,106.7,147.1,106.8,147.1,106.8z"
	pathData22 := "M172.9,136.8c0.8-1.8,1.6-2.6,2.7-3.1c1.2-0.5,1-1.4,0.5-2.1c-0.4-0.6-2.1-0.7-0.7-1.6    c0.7-0.5,1.8,0.5,1.7,1.4c-0.2,1.8,1.5,3.2,0.5,5.4C176.4,133.8,174.9,136.2,172.9,136.8z"
	pathData23 := "M164.9,164.4c1.3-2.5,3.9-3.1,5.4-5c0.3-0.4,0.9-0.6,1.3,0c0.3,0.5,0.2,0.9-0.4,1.2c-1.8,0.8-3,2.2-4,3.8    C166.1,166,165.8,163.2,164.9,164.4z"
	pathData24 := "M174.9,124.4c-3.2,1.1-3-2.2-4.9-3.7c1.6,0.3,2.6,0.6,3.8,0.7c0.7,0.1,1.4,0.5,1.5,1c0.1,1-1,0.3-1.3,0.7    c-1,1.1,1.3,0.6,0.8,1.5L174.9,124.4z"
	pathData25 := "M164.9,154.5c-0.4,3.6-0.8,4.1-4.3,5.2c1.4-1.6,3.1-2.4,2.2-4.9c-0.5-1.4,1.5,0.2,2.3-0.2    C165,154.7,164.9,154.5,164.9,154.5z"
	pathData26 := "M165.6,161.4c-1,1.4-1.7,2.3-3.2,2.4c-0.7,0-2.1,0.6-0.3,1.7c0.6,0.4-0.4,0.9-0.8,1.2c-0.5,0.3-1,0-1-0.6    c-0.1-1.2-0.1-2.8,1.1-2.8C162.9,163.1,163.7,161.5,165.6,161.4z"
	pathData27 := "M131.5,169.4c-1.5,1.8-2.6,4.5-6.3,2.6C127.7,171,129.6,170.2,131.5,169.4z"
	pathData28 := "M194.5,227.1c-1.7-2.9-2.9-4.9-4-6.9C192.8,221.1,194,223,194.5,227.1z"
	pathData29 := "M165.1,138.7c0.4,3,1.3,5.9-0.5,8.8C165.2,144.6,164.5,141.6,165.1,138.7z"
	pathData30 := "M118.2,154.5c-1.2-3.2,1.4-4.6,2.5-7.2C121.3,150.5,119.4,152.3,118.2,154.5z"
	pathData31 := "M152.3,165.8c-0.3-2.7,1.1-5.4,0.1-8.2C153.3,160.4,153.6,163.1,152.3,165.8z"
	pathData32 := "M155.4,165c0.6-2.8,0.3-5.6-0.4-8.3C157.2,159.3,156.3,162.2,155.4,165z"
	pathData33 := "M174,244.2c0.6-1.6,1.3-3.2,1.9-4.8c3.4,2.6-0.9,3.4-0.9,5.2C174.7,244.4,174.4,244.3,174,244.2z"
	pathData34 := "M179.7,245.4c2.1,1.5,1.5,3.7,2.1,5.9C180.4,250.2,180.4,250.2,179.7,245.4z"
	pathData35 := "M172.7,200.5c0.7,1.7,1.4,3.5,2.1,5.2c-2-1.1-2.2-3.1-3.1-4.7C172,200.9,172.4,200.7,172.7,200.5z"
	pathData36 := "M231.1,163.8c-1.4,1.2-3.7,1.3-4.3,3.5C226.7,164.2,228.3,163.2,231.1,163.8    C231.2,163.9,231.1,163.8,231.1,163.8z"
	pathData37 := "M229,123.4c-2.6,0.7-5.1,0.3-8.4-1.1C224,122,226.4,122.9,229,123.4C228.9,123.5,229,123.4,229,123.4z"
	pathData38 := "M180.1,256.8c0.4-1.9-2.3-2.5-1.8-4.6c2.4,0.7,2.8,3,3.2,4.8c0.2,0.9-0.9-0.5-1.5-0.3    C180,256.6,180.1,256.8,180.1,256.8z"
	pathData39 := "M188.5,223.5c2.3,0.8,4.2,2.3,3.8,3.4c-0.8,1.8-1.5-0.2-2.1-0.9C189.7,225.4,189.3,224.7,188.5,223.5z"
	pathData40 := "M156.8,93.1c-1.5,1.4-2,4.1-5.3,3.4C153.9,95.8,154.6,93.3,156.8,93.1z"
	pathData41 := "M192,229.6c0.6,1.6,0.7,2.8-0.4,5C190.9,232.4,190.4,231.1,192,229.6z"
	pathData42 := "M160.6,146.7c0.7-1.6,0.1-3.3,0.5-5.4C162,143.5,163.2,145.3,160.6,146.7z"
	pathData43 := "M174.8,124.5c0.8-0.3,1.8-1.3,1.6,0.6c0.2,1.4-0.3,1.9-1.7,1.7c-0.4,0.2-1,0.6-1.2,0.4    c-0.3-0.2-0.3-0.8,0.1-1.2c0.5-0.5,0.9-1.1,1.3-1.6C174.9,124.4,174.8,124.5,174.8,124.5z"
	pathData44 := "M174.7,126.8c0.8-0.3,1.4-0.9,1.7-1.7c0.7,2.2,0.7,2.2-2,4C174.5,128.3,174.6,127.5,174.7,126.8z"
	pathData45 := "M171.2,247.1c0.2,1.4,1.7,2.4,1.1,4c-0.1,0.2-0.9,0.4-0.9,0.3C170.5,250,171.2,248.5,171.2,247.1z"
	pathData46 := "M165.2,81.8c0.5,3.6-2.1,1-3,1.8C162.9,82.4,164.1,82.2,165.2,81.8L165.2,81.8z"
	pathData47 := "M168.7,101.3c-3.3,0.9-0.8-2.2-1.9-2.9C168.2,98.9,168.6,100,168.7,101.3L168.7,101.3z"
	pathData48 := "M144.8,164.5c-0.5,0.2-1,0.4-1.6,0.6c0.1-1.3-0.6-2.9,1.4-3.6C144.7,162.6,144.8,163.5,144.8,164.5z"
	pathData49 := "M149.1,165.1c0.1-1.4,0.3-2.8,0.5-4.2C151.6,162.4,150.1,163.8,149.1,165.1z"
	pathData50 := "M175.1,202.2c-0.3-1.1-0.5-2.3-0.8-3.4C176,199.6,175.1,201.1,175.1,202.2z"
	pathData51 := "M169.3,133.1c0.7-1,1.4-2.1,2.1-3.2C172.4,131.9,170.8,132.6,169.3,133.1z"
	pathData52 := "M166.5,106.9c-1-0.7-1.9-1.3-3-2.1C165.5,103.8,166,105.5,166.5,106.9z"
	pathData53 := "M188.8,247.5c0.5,1.1,0.9,2.2,1.5,3.6C187.9,250.6,189,248.8,188.8,247.5z"
	pathData54 := "M165,108.2c-2.1,0.1-2.3-1.3-3.5-1.8C163.5,105.6,163.8,107.3,165,108.2z"
	pathData55 := "M169.7,81.7c1.3-1,1.3-1,3.9-0.5C172.3,81.3,171,81.4,169.7,81.7L169.7,81.7z"
	pathData56 := "M124.4,152.1c-1.1-1.7,0.7-2.4,0.5-4C126.3,150,124.8,150.8,124.4,152.1z"
	pathData57 := "M172.8,245.4c0.4,1,1.4,1.5,1.1,2.6c-0.1,0.2-0.9,0.4-1,0.3C172.2,247.4,172.4,246.5,172.8,245.4z"
	pathData58 := "M159.1,131.2c1.9-0.9,3.2-0.4,4.5-0.3C162.3,131.6,161,131.4,159.1,131.2z"
	pathData59 := "M167,97.1c-1.1-0.2-2.2-0.4-3.3-0.6C164.9,96,166.1,95.8,167,97.1z"
	pathData60 := "M169.9,151.5c0.7,1.3,0.3,2.2-0.9,2.8C168.3,153.1,168.7,152.2,169.9,151.5z"
	pathData61 := "M187.4,249.8c0.4,1.3,1.8,2.1,0.9,3.7C187.8,252.5,186.8,251.7,187.4,249.8z"
	pathData62 := "M171.4,205.4c1.2,0.5,1.7,1.4,1.8,2.7C171.5,207.9,172,206.2,171.4,205.4z"
	pathData63 := "M194,259.9c0.4-1.3,0.7-2.4,1.4-4.3C195.6,257.7,196.1,258.9,194,259.9z"
	pathData64 := "M116.6,85.7c-0.2-1.6,1-1.9,1.7-2.9C119,84.8,117.2,84.9,116.6,85.7z"
	pathData65 := "M122.5,149.5c0.1-0.8,0.1-1.7,0.2-2.5c0.3,0.2,0.7,0.3,0.9,0.6C124.1,148.4,123.1,148.9,122.5,149.5z"
	pathData66 := "M135.5,146.4c0,0.5,0,1-0.7,1.1c-0.6,0.1-1-0.3-0.8-0.8c0.1-0.4,0.5-0.9,0.9-1    C135.8,145.3,135.3,146.2,135.5,146.4z"
	pathData67 := "M147.2,106.7c-0.7-0.1-1.4-0.1-2.2-0.2c0.4-0.4,0.7-1.1,1.2-1.2c1.2-0.2,0.9,0.8,0.9,1.5    C147.1,106.8,147.2,106.7,147.2,106.7z"
	pathData68 := "M126.8,84.4c-1.2,1.1-1.8,1.7-2.6,2.4C123.9,85.1,124.7,84.6,126.8,84.4z"
	pathData69 := "M169,91.2c0.4-1.1,1.1-1.7,2.3-1.2C170.9,91,169.9,91.1,169,91.2z"
	pathData70 := "M64.6,203.3c-1.1,0.8-2.2,0.5-3.4,0.3C62.3,202.6,63.4,202.9,64.6,203.3z"
	pathData71 := "M122.1,85.8c-0.4-1.4,0.3-1.9,1.4-2.1C124,85.1,122.4,85,122.1,85.8z"
	pathData72 := "M169.6,81.6c-1.4,1-2.9,0.7-4.4,0.2l0,0C166.7,81.1,168.2,81.2,169.6,81.6C169.7,81.7,169.6,81.6,169.6,81.6z    "
	pathData73 := "M163.1,85.6c1.8-0.7,2.6-0.1,3.2,0.8C165,87.3,164.4,85.7,163.1,85.6z"
	pathData74 := "M193,255.9c-0.3-0.8-1.5-1.3-0.7-2.4c0.1-0.2,0.6-0.4,0.6-0.3C193.5,254.1,193.3,255,193,255.9z"
	pathData75 := "M130.5,148.1c0.3-0.8,0.5-1.7,0.9-2.9C131.9,146.6,131.9,147.6,130.5,148.1z"
	pathData76 := "M173.3,119.1c-0.2,0.3-0.4,0.6-0.7,0.7c-0.7,0.1-1.1-0.4-1.4-0.9c-0.3-0.6,0.3-0.5,0.6-0.5    C172.5,118.3,173.1,118.3,173.3,119.1z"
	pathData77 := "M172.6,117.5c-0.1,0-0.4-0.2-0.7-0.3c0.4-0.4,0.8-0.9,1.3-1.1c0.6-0.2,0.8,0.4,0.7,0.9    C173.9,117.5,173.3,117.6,172.6,117.5z"
	pathData78 := "M129.4,231.6c0.9-1.2,1.8-2.4,2.8-3.7c0.5,0.8,0.1,1.4-0.1,2.2c-0.6,1.8-2.1,2.4-3.6,2.8    c-0.2,0.1-1-1-1.1-1.7c-0.3-1.7,1.8-2.5,1.7-4.2c0.2,0,0.6-0.1,0.7,0c0.2,0.3,0.5,0.7,0.4,0.9c-0.5,1.1-1,2.1-1.6,3.2    C128.8,231.2,129.1,231.4,129.4,231.6z"
	pathData79 := "M122.7,245.4c-0.5-0.7-1-1.4-1.6-2.3c1.9-0.5,2.3,2.4,4.3,1.4c-0.9,1.1-1.9,0.5-2.8,0.7    C122.5,245.3,122.7,245.4,122.7,245.4z"
	pathData80 := "M133.4,230c0.6-1.4,0.6-2.6,2-3.1C135.9,228.5,135.4,229.5,133.4,230z"
	pathData81 := "M122.9,231.5c-0.5-1.3-0.7-2.2,0-3.1c0.1-0.1,0.7,0.2,1,0.3C123.6,229.5,123.3,230.4,122.9,231.5z"
	pathData82 := "M164,219.7c-0.6,1.8-1.6,1.8-3.4,1.6C162,220.6,162.8,220.2,164,219.7z"
	pathData83 := "M158,224.9c0.6-0.9,1.3-1.5,2.5-1.1c0.1,0.1,0.2,0.1,0.3,0.2C159.8,224.2,159.3,225.8,158,224.9z"
	pathData84 := "M62.7,87.9c0.8,1.8,0.5,3.3-0.7,4.8C61.2,90.9,63.1,89.5,62.7,87.9z"
	pathData85 := "M65.1,93.3c-1.5-1.6-1.3-3.4-0.8-5.3C64.3,89.7,65.9,91.3,65.1,93.3C65.1,93.2,65.1,93.3,65.1,93.3z"
	pathData86 := "M150.2,86c0.2-1.4,0.1-3,2.2-3.1c2.1-0.2,2.8,1.7,2.9,2.8c0.1,2,1.5,3.8,0.8,5.8l0.1-0.1    c-2.2-0.4-5,0.8-6-2.4c0.1-0.1,0.3-0.2,0.4-0.3c0.7-0.3,1.8,0.2,2-0.6c0.2-1.2-1-1.5-2-1.8C150.5,86.2,150.4,86.1,150.2,86z"
	pathData87 := "M150.2,89c-2.2-1-2.2-2,0-3l-0.1-0.1c0.9,1.1,0.9,2.1,0,3.2L150.2,89z"
	pathData88 := "M150.1,89.1c0-1.1,0-2.1,0-3.2c2.9-2.4,2.9,1.7,4.5,2.2c0.4,0.1-0.7,1.5-1.8,1.7    C151.9,90,151.2,88.6,150.1,89.1z"
	canvas.Path(pathData0, "fill:"+lineColor)
	canvas.Path(pathData1, "fill:"+hexcode)
	canvas.Path(pathData2, "fill:"+hexcode)
	canvas.Path(pathData3, "fill:"+hexcode)
	canvas.Path(pathData4, "fill:"+hexcode)
	canvas.Path(pathData5, "fill:"+lineColor)
	canvas.Path(pathData6, "fill:"+hexcode)
	canvas.Path(pathData7, "fill:"+hexcode)
	canvas.Path(pathData8, "fill:"+hexcode)
	canvas.Path(pathData9, "fill:"+hexcode)
	canvas.Path(pathData10, "fill:"+hexcode)
	canvas.Path(pathData11, "fill:"+hexcode)
	canvas.Path(pathData12, "fill:"+lineColor)
	canvas.Path(pathData13, "fill:"+hexcode)
	canvas.Path(pathData14, "fill:"+lineColor)
	canvas.Path(pathData15, "fill:"+hexcode)
	canvas.Path(pathData16, "fill:"+hexcode)
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
	canvas.Path(pathData66, "fill:"+lineColor)
	canvas.Path(pathData67, "fill:"+lineColor)
	canvas.Path(pathData68, "fill:"+lineColor)
	canvas.Path(pathData69, "fill:"+lineColor)
	canvas.Path(pathData70, "fill:"+lineColor)
	canvas.Path(pathData71, "fill:"+lineColor)
	canvas.Path(pathData72, "fill:"+lineColor)
	canvas.Path(pathData73, "fill:"+lineColor)
	canvas.Path(pathData74, "fill:"+lineColor)
	canvas.Path(pathData75, "fill:"+lineColor)
	canvas.Path(pathData76, "fill:"+lineColor)
	canvas.Path(pathData77, "fill:"+lineColor)
	canvas.Path(pathData78, "fill:"+lineColor)
	canvas.Path(pathData79, "fill:"+lineColor)
	canvas.Path(pathData80, "fill:"+lineColor)
	canvas.Path(pathData81, "fill:"+lineColor)
	canvas.Path(pathData82, "fill:"+lineColor)
	canvas.Path(pathData83, "fill:"+lineColor)
	canvas.Path(pathData84, "fill:"+lineColor)
	canvas.Path(pathData85, "fill:"+lineColor)
	canvas.Path(pathData86, "fill:"+hexcode)
	canvas.Path(pathData87, "fill:"+hexcode)
	canvas.Path(pathData88, "fill:"+lineColor)
}
