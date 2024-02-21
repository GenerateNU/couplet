import React from "react";
import { Button } from "react-native-paper";
import { ReactionButtonProps } from "./ReactionButtonProps";

const ReactionButton = (props: ReactionButtonProps) => {
  return (
    <Button
      style={{
        padding: 5,
        paddingVertical: 20,
        flex: 1,
        alignSelf: "center" // This will center the button horizontally
      }}
      labelStyle={{ fontSize: 35, alignSelf: "flex-start" }}
      icon={props.icon}
      mode="elevated"
      onPress={() => props.handleReact(props.like)}
      children={undefined}
    />
  );
};

export default ReactionButton;
