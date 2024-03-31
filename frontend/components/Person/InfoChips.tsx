import React from "react";
import { Text, View, StyleSheet } from "react-native";
import { PillProps } from "./PersonProps";

export default function InfoChips({ items, textColor, backgroundColor }: PillProps){

    return (
        <View style={styles.container}>
            {items.map((item) => 
                <View style={{...styles.chipItemContainer, backgroundColor}}>
                    <Text style={{...styles.chipItemText, color: textColor}}>{item}</Text>
                </View>
            )}
        </View>
            
    )
}


const styles = StyleSheet.create({
    container: { 
        display: "flex", 
        flexDirection: "row",
        flexWrap: "wrap",
        justifyContent: "flex-start",
    },
    chipItemContainer: {
        border: "black", 
        borderRadius: 20,
        width: "auto",
        padding: 10,
        paddingLeft: 17, 
        paddingRight: 17, 
        margin: 5, 
    },
    chipItemText: {
        fontFamily: "DMSansBold",
        fontSize: 13,
    }
    
});