import { StyleSheet } from "react-native";
import { scaleWidth, scaleHeight } from "./ratio"; 

type StyleObject = {
  [key: string]: any; 
};

export default function scaleStyleSheet(rawStyles: { [key: string]: StyleObject }): any {
  const scaledStyles = Object.keys(rawStyles).reduce((acc: { [key: string]: StyleObject }, key) => {
    const style = rawStyles[key];
    const scaledStyle = Object.keys(style).reduce((accStyle: StyleObject, styleProp: string) => {
      const value = style[styleProp];

      const widthProperties = ['width', 'paddingHorizontal', 'marginHorizontal', 'left', 'right'];
      const heightProperties = ['height', 'paddingVertical', 'marginVertical', 'top', 'bottom', 'fontSize', 'lineHeight'];

      if (widthProperties.includes(styleProp) && typeof value === 'number') {
        accStyle[styleProp] = scaleWidth(value);
      } else if (heightProperties.includes(styleProp) && typeof value === 'number') {
        accStyle[styleProp] = scaleHeight(value);
      } else {
        accStyle[styleProp] = value;
      }

      return accStyle;
    }, {});

    acc[key] = scaledStyle;
    return acc;
  }, {});

  return StyleSheet.create(scaledStyles);
}
