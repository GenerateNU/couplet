export type ReactionButtonProps = {
  like: boolean;
  icon: string;
  handleReact: (like: boolean) => void;
};
