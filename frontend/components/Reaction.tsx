import React from "react";
import { View } from "react-native";
import ReactionButton from "./ReactionButton";

type ReactionProps = {
  handleReact: (like: boolean) => void;
};

const Reaction = (props: ReactionProps) => {
  return (
    <View
      style={{
        flexDirection: "row",
        justifyContent: "space-between",
        paddingTop: "10%"
      }}
    >
      <ReactionButton like={false} icon="cancel" handleReact={props.handleReact} />
      <ReactionButton like={true} icon="heart" handleReact={props.handleReact} />
    </View>
  );
};

export default Reaction;
