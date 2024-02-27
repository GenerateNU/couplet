export type ReactionButtonProps = {
  like: boolean;
  icon: string;
  label: string;
  handleReact: (like: boolean) => void;
};
