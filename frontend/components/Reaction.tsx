import React from "react";
import { View } from "react-native";
import ReactionButton from "./ReactionButton";

type ReactionProps = {
  handleReact: (like: boolean) => void;
};

function Reaction({ handleReact }: ReactionProps) {
  return (
    <View
      style={{
        flexDirection: "row",
        justifyContent: "space-between",
        paddingTop: "10%"
      }}
    >
      <ReactionButton like={false} icon="cancel" handleReact={handleReact} />
      <ReactionButton like icon="heart" handleReact={handleReact} />
    </View>
  );
}

export default Reaction;
