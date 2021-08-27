import { VFC } from "react";
import { IconButton } from "@chakra-ui/react";
import { ArrowBackIcon } from "@chakra-ui/icons";
import { useHistory } from "react-router-dom";

const GoBackBtn: VFC = () => {
  const history = useHistory();

  return (
    <IconButton
      aria-label="back to page"
      size="lg"
      icon={<ArrowBackIcon />}
      onClick={() => history.goBack()}
    />
  );
};

export default GoBackBtn;
