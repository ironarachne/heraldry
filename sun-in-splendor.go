package heraldry

import svg "github.com/ajstarks/svgo"

func renderSunInSplendorToSvg(canvas *svg.SVG, hexcode string, lineColor string) {
	pathData0 := "M207.1,205c8,11.9,15.7,23.5,23.4,34.9c0.8,1.2,2.5,2.4,1.1,3.7c-1.1,1-3.1,0.8-4.6-0.2c-2.4-1.5-4.9-3-7.4-4.5    c-12.7-8.2-25.3-16.5-38.1-24.8c-1.8,8.7,0.9,14.6,8.4,19c4.5,2.6,5.5,8.2,6.9,12.8c1.6,5.1,1.5,10.4,0.5,15.9    c-0.9,4.8,2.9,8.7,6.3,12c0.9,0.9,2.7,1.1,2.2,2.7c-0.5,1.7-2.2,1.5-3.7,1.6c-11,0.4-20.2-8.2-20-19.2c0.2-7.9-1.4-13.8-10.4-16.1    c-4.6-1.2-7.2-5.8-10.4-10.3c-4,15.1-7.7,29.4-11.5,43.7c-0.5,0-1,0-1.4,0c-5.8-17.9-8.7-36.6-13.6-54.7c-6.1,1.9-7,6.8-6.8,12.4    c0.5,10-2.8,18.1-10.8,24.7c-4.2,3.5-9.3,6.7-10.7,12.5c-0.8,3.1-0.6,6.5-0.8,10c-4.6-1-6.4-5-7.2-8.3c-2.1-8.1-2.3-15.8,4.7-22.9    c7.1-7.2,10.1-7.2,3.4-20.1c-1.7-3.3-0.6-8.1-0.8-12.7c-9.1,5.3-18.1,10.6-27.1,15.8c-3.1,1.8-6.3,3.3-9.1,5.3    c-1.7,1.2-3.5,3-5.1,1.4c-1.7-1.8,0.1-4,1.2-5.7c4.1-6.5,8.3-13,12.5-19.5c4.1-6.3,8.4-12.6,13.1-19.6c-5.8,1.1-10.9,0.5-14.6,5.8    c-6.5,9.4-15.9,12.9-27,9.3c-4.8-1.5-8,0.8-11.3,3.3c-1.8,1.4-3.9,4.3-5.7,2.6c-2-2-2.8-5.7-1.2-8.7c4-7.4,9.9-11.5,18.7-11.1    c8.7,0.3,9.1-0.6,12.9-8.7c2.7-5.7,5.6-11.8,13.6-14.2c-5.3-2.9-10.1-3.8-14.6-4.8c-8.5-1.9-16.6-5.3-25.3-6.4    c-1.3-0.2-3.5-0.8-3.7-2.7c-0.2-2.1,2.1-2.6,3.4-2.9c12.2-2.9,24.5-5.6,36.8-8.3c3.3-0.7,6.6-1.5,10-2c1.1-0.2,2.5-1,2.4-1.3    c-0.5-1-0.5-2.2-2-2.9c-4.9-2.3-10.1-1.9-15.1-1.7c-5.3,0.2-10-2-14-4.5c-4.1-2.7-6.3-7.7-8.7-12.2c-2.8-5.2-7.3-7.4-13-6.3    c-0.9,0.2-1.9,2-2.7,0.1c-0.4-0.9,0.2-1.8,0.9-2.6c5.3-6.8,14.2-7.4,20.6-1.5c3.8,3.5,6.6,8.3,12.3,9.5c1.5,0.3,3.1,0.9,4.3,0.4    c6.7-2.4,13.8-1.8,21.5-2c-7.6-10.6-14.7-20.6-21.8-30.5c-1.4-2-2.6-4.1-4.1-6c-0.9-1.2-2.7-2.4-1.6-3.7c1.3-1.6,3.3-0.5,4.9,0.4    c5.4,3.2,10.8,6.5,16.2,9.6c10.8,6.3,21.6,12.7,32.7,19.1c0.7-7-3.9-11-8.5-14.5c-6.1-4.8-9.3-10.7-9.4-18.3    c0-2.7,0.2-5.6,1.1-8.2c2.6-7.3,0.2-12.3-6.1-16.1c-1-0.6-1-1.4-0.1-2.3c1-1,2.1-1.4,3.5-1.4c7.9,0,17.7,6.6,17.4,16.4    c-0.1,2-0.3,4,0,6c0.7,4.1,1.3,8,6.3,9.8c5.8,2.2,11.1,5.4,15,10.9c3.4-14.6,6.7-28.8,10.2-42.9c0.2-0.9-0.5-3.5,2.1-3.1    c1.9,0.3,1.9,2.1,2.3,3.3c4.7,16,9.4,32,14,48c0.3,1.2,0.7,2,1.8,1.6c2.4-0.9,5.2-1.6,5.2-5c0.1-3.7,0.2-7.3,0-11    c-0.4-9.7,3.8-16.8,12.3-20.9c7.3-3.5,10.8-8.5,9.8-16.6c-0.1-0.4,0.2-0.9,0.3-1.3c7.7,2.7,10.5,13.2,6.9,20.2    c-2.6,5-7.7,7.7-9.9,12.8c-1,2.3-1.1,4.3-0.5,6.9c1.5,6,3.5,12.2,1,19.5c8-5.1,15.2-9.6,22.4-14.3c4.3-2.8,8.5-5.8,12.8-8.7    c1.3-0.9,2.8-1.9,4.2-0.4c0.9,1-0.2,2.1-0.8,3.1c-7.8,13.5-15.7,27-23.5,40.5c-1,1.6-1.8,3.4-2.6,5c6.8,1.8,11-0.7,15.2-6.1    c4.9-6.3,12.5-8.9,20.5-9.5c3-0.2,6-0.1,9,0c5.7,0.2,11.1-0.3,15-5.7c3.1,9.9-1.7,17.4-11.5,18.5c-3.4,0.4-6.7,1.2-10.2,1.3    c-5.1,0.1-7.9,4-9.8,8.2c-2.3,5.1-7.8,10-13.4,12.5c14,3.4,27.7,6.8,41.4,10.2c1.5,0.4,3.3,0.5,3.5,2.6c0.2,2.2-2,1.5-2.9,1.8    c-16.2,4.5-32.4,8.7-48.7,13.1c2.4,1.5,3.6,3.8,4.6,6.4c1.2,2.9,3.6,4.9,6.6,5.1c3.5,0.1,6.8,1,10.2,1.1    c8.9,0.3,15.1,4.7,17.9,12.6c2.6,7,6.6,9.8,13.8,8.5c1-0.2,2,0,3,0c-1,4.9-4.7,7.6-9.2,8.1c-7.6,0.7-13.6-0.8-18-7.9    c-3.6-5.8-9.7-6.3-16.1-4.2C215.2,207.9,216.8,208,207.1,205z"
	pathData1 := "M149.7,217c-18.6,0.9-33.4-6.9-45.5-20.4c-12-13.3-11.5-29.8-11.2-46.2c0.4-21.2,25.6-48.2,60.5-45.4    c27.5,2.2,52.2,23.2,51.4,56.9C204,193.9,182.8,216.2,149.7,217z"
	pathData2 := "M66.2,200.2c2.9-0.6,4.3-2.9,5.4-5.3c3.3-7.7,9.6-10.6,17.4-12c4.6-0.8,6.2,2.1,6.3,5c0.1,2.7-2.5,5.4-5.7,5    c-8-1.1-12.8,2.7-17,9c-3.3,5-12.9,9-19.5,6.2c-7.1-3-12.7,0.7-18.3,4.4c-0.8-4.3,1.2-8.3,5.2-10.8c4.3-2.6,8.6-4.1,13.8-3.8    c4.7,0.3,10.2-1.1,11.1-5.9c0.9-4.7,3.4-8,5.7-11.5c3.2-4.8,8.9-6.2,14.3-7.3c4-0.8,6.3,2.2,7.4,6.5c-2.9,0.5-5.9,0.9-8.7,1.7    c-6.7,2-11.1,6.8-13.1,13.1c-0.8,2.7-5.6,2.1-3.9,5.9c-4.4,1.1-8.4,3.8-11.5,3.3C57.8,203.8,62,201.9,66.2,200.2z"
	pathData3 := "M237.7,197.3c-2.3-3.4-9.3-5.4-13.6-3.9c-7.1,2.5-17.9-1.7-20.6-8.7c-0.9-2.3-0.7-6,2.8-7.8    c3.6-1.9,4.9,0.7,6.2,3.1c2,3.4,2.8,8.3,7.5,8.8c3.6,0.3,7.1,1.2,10.7,1.2c8.3-0.1,14.7,3.4,17.7,11.1c2.6,6.8,5.6,12.1,15.2,9.5    c-8.1,6.8-15.1,5.5-21.2-2.7c-5.3-7.2-14-9.1-21.7-4.8c-5,2.8-14.3,0.5-17.7-4.3c-1.9-2.7-4.9-5.1-2-9c1.2-1.7,2.5-2.6,3.7-0.5    c4.2,7.4,11.6,7,18,6.4c5.4-0.6,10.2,0.5,15.1,1.5c0.3,0.1,0.5,0.2,0.6,0.2C238.3,197.3,238,197.3,237.7,197.3z"
	pathData4 := "M121.5,214.7c-5.1,5.7-7,12.4-4.4,19.9c1.2,3.6,1.4,6.9-0.5,10.3c-0.8,1.4-1.1,3.1-0.9,4.9    c2.7-4.1,4.6-8.4,4-13.7c-0.5-3.9-2.3-7.6-0.8-11.8c1-2.8,2.7-5,4-7.6c0.7-1.4,8.4-0.7,9.1,0.7c0.4,0.8,0.4,2.2-0.7,2.7    c-6.6,2.6-5.2,8.3-5.2,13.4c0.1,10.1-3.6,18.2-12.1,24.1c-2.8,2-5.2,4.7-7.4,7.4c-2.8,3.4-4.3,7.2-2.2,12.7    c-6.1-7.6-5.2-20.3,0.6-25.6c4.1-3.8,9.6-8,8.2-15.4c-0.2-1.3-0.2-2.7-1.1-3.6c-4-4.8-3-10.5-3-16c0.1-3.4,7-6.7,9.8-4.7    C119.8,213.1,121.3,212.7,121.5,214.7z"
	pathData5 := "M179.6,236.4c-1.6-4.4-6.3-6.6-7.4-11.5c-1.4-6.5-0.2-9.8,5.7-11.5c0.1,5.8-0.6,11.8,4,16.4    c0.6,0.6,0.9,1.7,1.8,2.2c10.1,5.5,12.3,15.2,10.9,24.9c-1.1,7.4,1.8,12.4,5.2,17.8c-7.4,0.3-14.8-7.4-14.7-15.9    c0-3.5-0.4-6.9-1.2-10.2c-1.2-4.5-5.4-6.2-9.4-7.6c-8.8-3.1-13.7-13.6-10.4-22.2c1-2.5,3.7-1.5,6.1-3.8    c-1.1,9.7,0.7,17.3,9.6,21.2c3.5,0.1,4.5,3.4,6.6,5.2c4.9,7.3,2.4,15.4,2.5,19.7c-0.2-4.3,2.6-12.6-2.7-19.9    C184.1,239.6,183.3,236.1,179.6,236.4z"
	pathData6 := "M105.7,58.6c0.8,6.9-0.3,14,0.5,21c0.5,4.6,3.9,8.1,8.2,10.2c3.3,1.6,6.3,3.4,8.6,6.6c4.7,6.5,4,9.8-3.6,12.7    c-0.7-7.8-6-12.2-11.5-17c-5-4.5-8.2-10.2-6.2-18.1c1.9-7.8,3.1-16.3-5.5-22.8c9.4,0.9,14.8,6.6,14.8,14.6c0,2.2-0.2,4.4,0.1,6.5    c0.9,6,3.9,10.3,9.8,12.3c4.4,1.5,7.6,4.9,11.1,7.8c3.5,2.9,2.9,6.3,2.9,9.9c0,2.5-2.2,2.4-3.8,2.8c-1.8,0.4-1.9-0.8-2.2-2.3    c-1.4-6.4-5.7-10.8-11.4-13.5c-6.2-2.9-10.6-9.1-10.5-16c0-2.4,0.6-4.8,0.4-7.2C107.2,63.5,106.3,61.1,105.7,58.6z"
	pathData7 := "M222.1,132.9c-1.3,0.2-2.5,0-2.8,0.4c-3.7,5.8-9.6,4.5-14.9,4.5c-2.7,0-3.6-2.7-3.3-4.6    c0.3-1.8,2.3-2.9,4.5-2c5.5,2.1,10-0.1,12.7-4.5c5.2-8.6,13.8-10.2,22.3-11.7c6.6-1.1,13.4,0.5,20.3-1.2c-1.3,5.4-4.3,8.2-9.6,9.1    c-4.1,0.7-8.3,1.5-12.4,2c-5.2,0.7-7.9,4.7-10,9.1c-4.1,8.2-12.7,13.4-20.5,12.9c-2-0.1-2.8-0.7-2.4-2.6c0.2-1.4-1-3.6,2-3.5    C213.9,141.2,218.2,138.2,222.1,132.9z"
	pathData8 := "M35.1,114.8c5.8-5.4,10-3.6,14.1-0.1c1.9,1.6,3.6,3.5,5.3,5.3c6,6.3,13.4,6.3,20.8,3.8    c5.8-1.9,17.6-0.3,20.4,5.6c0.5,1,2.4,2.5,0.8,3.3c-1.3,0.6-2.3,3.2-4.5,1.3c-0.9-0.7-2.6-0.6-3.5-1.3c-6.5-4.8-13.3-2.6-20.3-1.6    c-4,0.6-8.1-1.2-11.2-4.4c-1.8-1.9-3.6-3.8-5.8-6.1c3.3,9.3,11.4,13.3,20.4,11.9c6.5-1,13.4-1.6,18.9,4.2c2.1,2.2,2.2,4,0.6,5.7    c-1.8,1.9-3.8,0.2-4.9-1.1c-2.8-3.5-6.3-4.5-10.7-4.2c-4.7,0.3-9.5,0.5-14.2-0.9c-5.1-1.6-9-4.9-11.1-9.6    C46.9,119.5,45.4,114.9,35.1,114.8z"
	pathData9 := "M148.6,55.6c4.1,13.3,8.2,26.5,12.2,39.8c1.8,6.1-0.8,8.5-7.1,7.6c-4.8-0.8-6-2.7-5.8-7.5    c0.5-13.3,0.2-26.6,0.2-39.9C148.2,55.5,148.4,55.6,148.6,55.6z"
	pathData10 := "M193.9,62.1c3.1-2,2.7-5.9,4.6-8.7c2.8,6.4-1,10.7-4.6,15.2c-1.6,2-3.5,3.8-5.2,5.9c-5.1,6.3-0.5,12,0.4,17.8    c1.2,7.2-1.9,15.1-8.2,18.4c-2.1,1.1-2.5-1-3.5-1.7c-1-0.7-1.7-1.3-0.3-2.6c6.5-6.3,5.5-14.1,4-22c-0.7-3.9,1.3-7.6,2-11.4    c0.1-0.2,0.2-0.4,0.3-0.6c0.7-0.8,1.4-1.6,2-2.5c0.6-0.8,1.4-1.6,2.1-2.4c0.9-0.9,1.8-1.7,2.7-2.6c0.3-0.3,0.6-0.5,1-0.7    c0.8-0.5,1.6-1.2,2.3-1.9C193.6,62.2,193.8,62.2,193.9,62.1z"
	pathData11 := "M188,67.9c-0.4,1-1.1,1.8-2.2,2.2c-1,1-2,2-3,3c0,0,0.1-0.1,0.1-0.1c-3.7,2.1-5,7.2-3.6,13.1    c1.6,6.5,0.2,16.9-6,20c-3,1.5-5.8,0-7.1-1.7c-1.7-2.1,1.2-3.2,2.9-4c3.5-1.6,4.9-4.1,4.9-8.1c0-5-0.3-10.1,0.6-15.2    c1.1-6.4,5.1-9.9,9.9-13c1.9-1.2,4-1.9,5.9-3.2c4.1-2.8,5-7,4.4-12.8c4.5,6-1,9.7-1,14c0,0,0.1-0.1,0.1-0.1    c-0.8,1.3-1.8,2.4-3.1,3.1l0.2-0.1C189.8,65.7,188.8,66.6,188,67.9z"
	pathData12 := "M223.2,88.8c-6.8,11.6-13.5,23.3-20.4,34.8c-3.9,6.4-4.3,6-8.4-0.6c-2.3-3.7-1.8-5.5,1.2-8.3    c9.2-8.7,17.9-17.8,26.9-26.8C222.7,88.3,222.9,88.5,223.2,88.8z"
	pathData13 := "M40.6,158.4c16.4-3.7,32.8-7.4,49.2-11.1C90.4,159,90.4,159,79,159c-12.8,0-25.6,0-38.4,0    C40.6,158.8,40.6,158.6,40.6,158.4z"
	pathData14 := "M226.8,239.6c-12.1-12.1-24.2-24.2-36.4-36.4c-0.1,0.2,0-0.1,0.2-0.3c6.7-6.8,7-7,12.8,1.2    c8.2,11.5,15.8,23.4,23.7,35.1C227,239.3,226.9,239.5,226.8,239.6z"
	pathData15 := "M73.6,88c12.5,7.4,25.1,14.7,37.6,22.1c1.8,1.1,2.4,2.2,0.5,3.8c-0.5,0.4-0.9,0.9-1.4,1.4    c-5.3,5.3-5.3,5.3-10.4,0.3c-9-8.9-17.9-17.8-26.9-26.8C73.2,88.5,73.4,88.2,73.6,88z"
	pathData16 := "M73.8,231.7c9.9-9.9,19.4-19.3,28.8-28.9c1.5-1.5,2.4-1.8,3.7-0.1c0.5,0.6,1.4,0.7,2,1.2    c5.6,4.6,5.4,5.3-0.8,8.9c-10,5.8-19.8,11.8-29.7,17.7C76.9,231.1,75.9,232.2,73.8,231.7z"
	pathData17 := "M69.7,232.2c6-9.2,12.1-18.4,18.1-27.6c2.6-3.9,5.1-7.8,7.7-11.6c0.4-0.5,0.5-2.2,1.1-1.6    c2.2,2.4,4,5,5.9,7.6c1.4,2-0.8,2.6-1.7,3.5c-10,10.2-20.1,20.3-30.2,30.4C70.4,232.7,70,232.4,69.7,232.2z"
	pathData18 := "M252.4,160c-14.1,0-28.2-0.1-42.3,0.1c-5.5,0.1-2.5-3.8-3-5.9c-0.5-2.3-0.5-4.3,3.2-4.2    c9.2,0.2,17.6,3.8,26.4,5.5c5.3,1,10.5,2.7,15.7,4C252.4,159.6,252.4,159.8,252.4,160z"
	pathData19 := "M47.6,161c13.3,0,26.5,0,39.8,0c1.7,0,2.6,0.2,2.7,2.3c0.6,7.3-2.9,9.7-10.2,7.2    C69.4,166.8,58.5,164.2,47.6,161z"
	pathData20 := "M247.5,162.9c-13.1,3.4-26.3,6.7-39.4,10c-0.5,0.1-1.7-0.1-1.8-0.4c-1-3.3,1.2-6.3,0.8-9.6    c-0.2-1.6,2.1-0.9,3.3-0.9c12.3-0.1,24.6,0,36.9,0C247.4,162.3,247.5,162.6,247.5,162.9z"
	pathData21 := "M150,261.4c0-12.8,0.1-25.6-0.1-38.4c0-2.4,0.8-3,3.1-3c8.8-0.2,8.8-0.1,6.5,8.1c-3.1,11.1-6,22.2-8.9,33.3    C150.4,261.4,150.2,261.4,150,261.4z"
	pathData22 := "M98.3,127.1c-11.9-12.4-19.4-27.6-31-40.5c1.9-0.4,2.4,0.8,3,1.4c9.4,9.4,18.9,18.8,28.3,28.2    C103.9,121.4,103.9,121.6,98.3,127.1z"
	pathData23 := "M147.4,262.3c-2.2-9.7-4.3-19.3-6.7-29c-0.9-3.4-1.1-7-2.5-10.3c-0.6-1.5-0.5-3.4,2.2-3.1    c1.6,0.2,3.3,0.2,5,0c2.3-0.2,2.7,0.7,2.7,2.8c-0.1,13.3,0,26.5,0,39.8C147.8,262.5,147.6,262.4,147.4,262.3z"
	pathData24 := "M217.3,234c-11.2-7-22.5-14-33.7-21.1c-1.4-0.9-2.5-3-1.2-3.5c2.5-0.9,3.4-5.3,7-3.4c1.1,0.6,2.3,1.3,3.2,2.2    c8.5,8.4,16.9,16.9,25.4,25.3C217.8,233.7,217.5,233.8,217.3,234z"
	pathData25 := "M183.2,112.4c11.1-8.7,23.5-15.5,35.7-22.8c-5.7,5.7-11.5,11.5-17.2,17.2c-2,2-4,4-6,6    C189.2,119.3,189.2,119.3,183.2,112.4z"
	pathData26 := "M146,61.6c0,12.8-0.1,25.6,0.1,38.4c0.1,3.7-2.1,2.9-4.1,3c-2.1,0.1-4,0.5-4.1-3c-0.3-10.3,3.7-19.7,5.4-29.5    C143.8,67.5,145.2,64.7,146,61.6z"
	pathData27 := "M161.9,164.1c3.6,0.1,4.8,2.9,4.8,5.6c0,3.4-2.7,3.7-5.5,3.4c-3-0.3-5.7,0.2-8.2,2.3c-1.8,1.5-3.9,2.1-6.7,0.9    c-2.6-1.2-5.4-2.8-9.1-1.5c-3,1.1-4.6-2.7-6.3-4.7c-1.1-1.3,0.1-3.5,1.5-4.9c1.8-1.9,3.4-3.8,5.3-5.8c1,3.1-1.6,4-2.9,5.4    c-1.5,1.6-2.4,2.9-0.9,5.1c0.8,1.2,1.2,2.3,2.8,2.2c2.2-0.2,1.3-2,1.3-3.1c-0.2-3,0.3-5.8,2.1-8.3c2.5-3.6,0.3-7.3-0.2-10.8    c-0.2-1.5-1.5-0.2-2.3,0.1c-4.1,1.4-8.1,3-12.3,4c-2.2,0.5-4.9,0.6-6.8-0.2c-3.6-1.6-7.3-1.5-11.1-1.6c-0.7,0-1.4,0.2-2,0    c-2.5-0.8-4.1-2.8-5.2-4.9c-0.7-1.3,0-2.8,2-1.4c1,0.7,1.3,1.9,2.4,2.7c3.3,2.5,5.4,3.1,7.2-1.6c1.5-3.9,11.7-7,16.9-5.5    c1.5,0.4,2.9,1.2,4.4,1.7c1.3,0.4,3.2,2.3,3.6,0.9c0.5-1.6,2.3-3.7-0.1-6c-3.3-3.2-6.6-4.7-11.4-4.3c-1.8,0.1-3.9-1.2-6.5-0.9    c-5.8,0.7-10.3,3.2-14.4,7.3c0.1-3.7,5-7.7,10.4-9c6.7-1.7,13-2.4,19.4,1.4c1.5,0.9,5.9,4.2,6-1.8c4.3,6.5,0.9,13.7,1.9,20.5    c0.7,4.7,1.4,9.1-0.9,13.7c-1,1.9-0.4,4.8,1.7,6.1c2.9,1.7,5.9,3.2,9.6,1.3c2.8-1.5,3.5-3.6,3.5-6.3c0.1-3.5-2.8-6.3-2-9.9    c1.3-6.5,0.5-12.7-1.9-18.9c-0.2-0.6-0.6-2.2,1.2-1.7l-0.1-0.1c2.2,5.3,6.1,0.9,7.4,0.2c8.5-4.6,16.9-5.6,25.8-1.6    c0.9,0.4,2.1,0.3,4.7,0.6c-5.9,0.6-2.6,2.5-1.7,4.1c-12.6-8.2-23.2-5.5-32.7,4.3c2.2,1.2,4.7,4.4,7-0.9c1.2-2.7,10.5-5.6,13.2-4.8    c1.3,0.4,1.9,1.2,1.9,2.5c-7.9,1.1-7.9,1.1-11.9,4.7c5-2.1,9.2-1.4,13.1-1.8c3.7-0.4,5.2,2.6,5.3,3c0.6,3.2,2.4,4.2,5.3,3.6    c3-0.6,4.8-2.3,4.4-6.2c2.6,5,1,8-4.7,8.1c-4.1,0.1-7.6,1.4-11.3,2.6c-6.1,1.9-11.8-0.3-17.7-1.6c-0.3-0.1-0.5-0.4-0.8-0.6    c-1.2-1.1-0.2-3.4-2.3-4c-0.2,0-0.7,0.4-0.7,0.7c-0.3,2.8-3.5,5.5-0.1,8.3c0.2,0.2,0.2,0.6,0.3,0.9c0,0.2,0,0.4,0,0.6    c0,1,0.3,2,0.7,3c0.2,0.4,0.3,0.8,0.4,1.3c0.7,2.3,0.6,4.9,2.8,7c1.2-1.5-0.1-2.9-0.1-4.1C161.8,164.5,161.9,164.3,161.9,164.1z"
	pathData28 := "M172.5,189.1c-1.5-1.6-2.4-2.5-3.5-3.7c0.7-0.2,1.1-0.5,1.4-0.4c4.4,1,3.7,4.3,3.3,7.4c-0.1,0.6-1.1,1.4-1.7,1.4    c-1.3,0.1-0.9-0.8-0.7-1.7c0.3-2.2-1.3-1.1-1.9-0.9c-2.9,1.2-5,3-6.5,5.9c-3.5,6.8-14.3,8.2-20.1,3.1c-3-2.6-4.9-6.2-8.7-8.1    c-1.9-1-3.6-1.2-5.6-1.2c-2.4,0-5.9-0.8-0.9,5c-6.1-4.8-5.9-6.4-1.2-10.9c1.8,3.1,6.2,4.3,9.6,1.7c4.1-3,8.1-4,12.7-1.6    c0.3,0.2,0.7,0.6,0.8,0.6c7.3-5.1,12.8,1.1,18.9,2.9C169.3,189,170.3,188.9,172.5,189.1z"
	pathData29 := "M135.7,141.1c-6.5-1.1-12.8-4.5-19.2,0.3C122.5,136,128.2,135.9,135.7,141.1z"
	pathData30 := "M94.8,165.7c2.1,3.4,4.4,6.8,2.8,11C96.6,172.9,95.7,169.3,94.8,165.7z"
	pathData31 := "M157.9,157.1c2.6,1.6,2.9,4.5,4,7c0,0,0.1-0.2,0.1-0.2c-3.2,1.5-2.8-1.9-4.1-3l0.1,0.2c-0.8-1.4-0.9-2.8,0-4.2    L157.9,157.1z"
	pathData32 := ""
	pathData33 := "M103.7,180c1.2,2.3,2.4,4.7,3.8,7.5C104.9,186,104.9,186,103.7,180z"
	pathData34 := "M177.6,205.4c-2.3,0.8-3.2,4.6-6.6,3.2C172.9,206.6,175.2,205.9,177.6,205.4z"
	pathData35 := "M193.6,141.6c-0.2-2-3.5-1.2-3.2-3.4c0.1-0.7,0.7-0.8,1.8-0.5C194.8,138.4,193.9,140,193.6,141.6z"
	pathData36 := "M137.9,109.5c-2.6,0.7-4.8,2.4-7.8,1C132.7,109.8,135,108.1,137.9,109.5z"
	pathData37 := "M99.4,179c0.3,0.8,0.6,1.1,0.5,1.3c-0.4,1.9,2.6,2.3,1.8,4.2c-0.1,0.2-1.2,0.4-1.2,0.3C99.5,183.2,98.4,181.5,99.4,179z"
	pathData38 := "M185.5,180.9c1,2,0.3,3.2-1.2,4c-0.3,0.2-1.2-0.2-1.2-0.4C182.9,182.6,184.7,182.3,185.5,180.9z"
	pathData39 := "M100.5,173.1c-0.6-2.5-1.1-4.4-1.8-7.1C101.3,168.3,101.4,170.4,100.5,173.1z"
	pathData40 := "M200.7,154.6c0,2,0,4,0,7.1C199.7,158.6,199.9,156.6,200.7,154.6z"
	pathData41 := "M180.9,180.8c-0.8,1.8-0.8,3.4-2.6,4C177.6,182.7,178.2,181.5,180.9,180.8z"
	pathData42 := "M198.5,169.5c-0.6-1.8-0.8-3.4,1-4.3C201.1,167.2,198.7,168.2,198.5,169.5z"
	pathData43 := "M118.4,204.9c-2.7,0.1-3.2-1.5-4.7-2.9C116.3,202,116.8,203.6,118.4,204.9z"
	pathData44 := "M188.7,173.7c0.3,2.1,0.7,3.2-0.2,4.2c-0.2,0.2-1.1,0-1.2-0.2C186.7,176.2,188.1,175.5,188.7,173.7z"
	pathData45 := "M114.5,189.3c-2.1-1.3-3.2-2-3.5-3.8C112.8,185.6,113.5,186.8,114.5,189.3z"
	pathData46 := "M121.5,118.2c-0.4,1.9-2.1,1.7-3.6,3.1C118.3,118.2,120.4,118.8,121.5,118.2z"
	pathData47 := "M162.4,114.5c1.7-0.8,2.8-0.3,3.4,1.1C164.1,117,163.7,114,162.4,114.5z"
	pathData48 := "M198,174.2c0,1.1,0,1.9-1.1,1.7c-0.3,0-0.7-0.6-0.8-0.9c0-0.9,0.7-1.4,1.4-1.7C197.6,173.3,197.9,174,198,174.2z"
	pathData49 := "M153.3,135.3c-0.8-1.4-0.1-2.4,1.1-3.1c0.2-0.1,1,0.5,1,0.8c0,1.5-1.7,1.1-2.3,2.1C153.2,135.2,153.3,135.3,153.3,135.3z"
	pathData50 := "M66.2,200.2c-4.2,3.5-9.2,4.4-14.5,3.3c5.1-0.2,9.5-3.2,14.7-3.2C66.3,200.3,66.2,200.2,66.2,200.2z"
	pathData51 := "M50.3,202.6c-2.7,0.3-4.5,2.4-6.9,1C45.2,202,47.3,201.5,50.3,202.6z"
	pathData52 := "M246.1,204.8c2.3,2.3,4.8,3.9,6.1,6.8C248.2,210.6,247.7,210,246.1,204.8z"
	pathData53 := "M237.7,197.3c2-0.6,3.7-0.5,4.4,2.3c-1.9-0.5-3.4-0.9-4.3-2.5C237.8,197.2,237.7,197.3,237.7,197.3z"
	pathData54 := "M108.7,257.3c-0.7,1.1-1.4,2.2-2.2,3.3c-0.9-0.7-0.7-1.6-0.2-2.4c0.4-0.6,1.1-1,1.7-1.5    C108.2,257,108.5,257.2,108.7,257.3z"
	pathData55 := "M185.7,241.3c3.8,2,3.9,5.7,4.3,9.3c0.6,6-1.6,12.2,2.5,18.3c-4.7-3-5.7-4.3-4.3-9.9c1.4-5.9,0.8-11.1-2.1-16.2    c-0.2-0.4-0.1-1-0.2-1.5L185.7,241.3z"
	pathData56 := "M179.6,236.4c4.1-1,5.8,1.1,6.1,4.9c0,0,0.1-0.2,0.1-0.1c-3.3,0-3-4.6-6.1-4.9C179.7,236.3,179.6,236.4,179.6,236.4z"
	pathData57 := "M222.3,132.3c0.4-5.4,5.7-11.1,10-11.3c0.7,0,1.4,0,1.6,0.7c0.1,0.3-0.4,1.1-0.6,1.1c-2.7-0.3-4.5,1.7-6.2,3.1    C225.2,127.5,224.1,129.9,222.3,132.3z"
	pathData58 := "M182.9,73.1c0.1-1.9,0.7-3.2,3-3C186.2,72.4,184.8,73,182.9,73.1z"
	pathData59 := "M190.9,65.1c0.7-1.4,1.7-2.5,3.1-3.1C194.1,64,193,65.1,190.9,65.1z"
	pathData60 := "M188,67.9c-0.2-2.2,1.1-2.9,3-3C190.4,66.3,189.5,67.4,188,67.9z"
	pathData61 := "M178.9,144.9c2,1,1.5,3.9,4,4.6c-1.7,1-3.4,0.7-4.9,1.8c-2.6,2-7.2,0.1-7.8-3.1c-0.9-4.5,3.1-3.1,5.2-4.4    c0.6,1.9-1.9,4.3,0.8,5.1C179.4,149.8,177,145.7,178.9,144.9z"
	pathData62 := "M116.1,147.6c5.1-5.4,7.1-5.7,10.8-2.6c1.7,1.4,1.4,2.8,0.4,4.1C125.2,152,121.8,151.5,116.1,147.6z"
	pathData63 := "M157.9,160.9c2.1,0,1.5,3.8,4.1,3c1.3,2,2.8,4,1.7,6.6c-0.1,0.3-1.2,0.6-1.7,0.4c-2.5-0.9-2.8-3-3.3-5.3    C158.3,164,158.2,162.5,157.9,160.9z"
	pathData64 := "M130.5,150.2c-0.9-2.3,0.1-4.3-1-6.1c1.9,1.2,5.3,0.6,5.4,3.6C134.9,150.5,131.4,147.7,130.5,150.2z"
	pathData65 := "M167.6,152c-1.5-1.9-2.7-2.8-2.9-4c-0.1-0.9,1.5-1.4,2.4-0.9C168.9,148,167.8,149.5,167.6,152z"
	pathData66 := "M158.1,156.9c0,1.4,0,2.8,0,4.2C157,159.7,155,158.3,158.1,156.9z"
	pathData67 := "M162.7,191.9c-2,5.4-4.3,8.3-8.3,9c-5.7,1.1-10.2-1.3-15-7.9C147.2,193.3,154.4,194.7,162.7,191.9z"
	pathData68 := "M135.3,189.8c3.7-1.5,6.4-4.5,8.4-3.5c4.1,2.1,7.2,1.9,11.5,0.1c2.4-1,6.3,1.8,10.5,3.3    c-5.6,0.9-10.4-0.9-14.6,2.2c-0.6,0.4-2.1,0.3-2.8-0.2C144.7,189,140.4,190.4,135.3,189.8z"
	pathData69 := "M122.7,145c1.5-0.1,2.2,0,2.3,1.1c0.1,1.3-0.5,2.2-1.6,2.7c-1.1,0.4-1.4-0.4-1.3-1.3C122.1,146.5,121.4,145.1,122.7,145z"
	canvas.Translate(10, 10)
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
	canvas.Path(pathData14, "fill:"+hexcode)
	canvas.Path(pathData15, "fill:"+hexcode)
	canvas.Path(pathData16, "fill:"+hexcode)
	canvas.Path(pathData17, "fill:"+hexcode)
	canvas.Path(pathData18, "fill:"+hexcode)
	canvas.Path(pathData19, "fill:"+hexcode)
	canvas.Path(pathData20, "fill:"+hexcode)
	canvas.Path(pathData21, "fill:"+hexcode)
	canvas.Path(pathData22, "fill:"+hexcode)
	canvas.Path(pathData23, "fill:"+hexcode)
	canvas.Path(pathData24, "fill:"+hexcode)
	canvas.Path(pathData25, "fill:"+hexcode)
	canvas.Path(pathData26, "fill:"+hexcode)
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
	canvas.Path(pathData61, "fill:"+hexcode)
	canvas.Path(pathData62, "fill:"+hexcode)
	canvas.Path(pathData63, "fill:"+hexcode)
	canvas.Path(pathData64, "fill:"+hexcode)
	canvas.Path(pathData65, "fill:"+hexcode)
	canvas.Path(pathData66, "fill:"+hexcode)
	canvas.Path(pathData67, "fill:"+hexcode)
	canvas.Path(pathData68, "fill:"+hexcode)
	canvas.Path(pathData69, "fill:"+lineColor)
	canvas.Gend()
}
