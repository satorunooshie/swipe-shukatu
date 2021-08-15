import { VFC } from "react";
import {Lottie} from "@crello/react-lottie";
import animationData from "../../lib/890-loading-animation.json";

const Loading: VFC = () => {
  return <Lottie config={{animationData: animationData, loop: true}} width="300px" height="300px"/>;
};

export default Loading;
