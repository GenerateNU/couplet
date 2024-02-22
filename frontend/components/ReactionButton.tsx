import React from "react";
import { Button } from "react-native-paper";
import { ReactionButtonProps } from "./ReactionButtonProps";

/* eslint-disable react/no-children-prop */

function ReactionButton({ like, icon, label, handleReact }: ReactionButtonProps) {
  return (
    <Button
      style={{
        padding: 5,
        paddingVertical: 20,
        flex: 1,
        alignSelf: "center" // This will center the button horizontally
      }}
      labelStyle={{ fontSize: 35, alignSelf: "flex-start" }}
      icon={icon}
      mode="elevated"
      onPress={() => handleReact(like)}
    >
      {label}
    </Button>
  );
}

export default ReactionButton;
