import { VFC } from "react";
import {
  Flex,
  IconButton,
  Input,
  Button,
  FormControl,
  FormErrorMessage,
} from "@chakra-ui/react";
import { MAIN_COLOR } from "../../constants/MainColor";
import { AttachmentIcon } from "@chakra-ui/icons";
import { FaTelegramPlane } from "react-icons/fa";
import { useForm } from "react-hook-form";

type FormData = {
  message: string;
};

const ChatForm: VFC = () => {
  const {
    handleSubmit,
    register,
    setValue,
    formState: { errors, isSubmitting },
  } = useForm<FormData>();

  const onSubmit = handleSubmit((data) => {
    // TODO: API call
    return new Promise((resolve) => {
      setTimeout(() => {
        alert(JSON.stringify(data));
        setValue("message", "");
        resolve(1);
      }, 2000);
    });
  });

  return (
    <form onSubmit={onSubmit}>
      <FormControl
        id="inputText2"
        isRequired
        isInvalid={errors.message ? true : false}
        textAlign="center"
      >
        <Flex
          pos="sticky"
          bottom="0"
          zIndex="2"
          w="full"
          minH={"50px"}
          align={"center"}
          justify="space-between"
          bg="white"
          pt="4"
          borderTop="1px solid"
          borderColor="gray.200"
        >
          <IconButton
            variant="ghost"
            aria-label="attach image"
            icon={<AttachmentIcon />}
          />
          <Input
            variant="filled"
            placeholder="テキストを入力してください"
            aria-label="input message"
            ml="2"
            {...register("message", {
              required: "This is required",
              maxLength: { value: 140, message: "140字以内で入力してください" },
            })}
          />
          <Button
            colorScheme={MAIN_COLOR}
            loading="false"
            rightIcon={<FaTelegramPlane />}
            aria-label="send message"
            ml="2"
            isLoading={isSubmitting}
            loadingText="送信中"
            type="submit"
          >
            送信
          </Button>
        </Flex>
        <FormErrorMessage>
          {errors.message && errors.message.message}
        </FormErrorMessage>
      </FormControl>
    </form>
  );
};

export default ChatForm;
