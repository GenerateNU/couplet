import React from "react";
import { Button } from "react-native-paper";
import { ReactionButtonProps } from "./ReactionButtonProps";

/* eslint-disable react/no-children-prop */

function ReactionButton({ like, icon, label, handleReact }: ReactionButtonProps) {
  return (
    <Button
      style={{
        paddingVertical: 10,
        marginHorizontal: 10,
        borderRadius: 30,
        backgroundColor: like ? "black" : "white",
        flex: 1,
        alignSelf: "center" // This will center the button horizontally
      }}
      labelStyle={{ fontSize: 20, alignSelf: "flex-start", color: like ? "white" : "black" }}
      icon={icon}
      mode="elevated"
      onPress={() => handleReact(like)}
    >
      {label}
    </Button>
  );
}

export default ReactionButton;
