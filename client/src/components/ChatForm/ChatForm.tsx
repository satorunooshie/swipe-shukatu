import { VFC, useState } from "react";
import {
  Flex,
  IconButton,
  Input,
  Button,
  FormControl,
  FormErrorMessage,
  Popover,
  PopoverTrigger,
  PopoverContent,
  PopoverHeader,
  PopoverBody,
  Switch,
  PopoverArrow,
  PopoverCloseButton,
  useDisclosure,
  FormLabel,
} from "@chakra-ui/react";
import { MAIN_COLOR } from "../../constants/MainColor";
import { AttachmentIcon, BellIcon } from "@chakra-ui/icons";
import { FaTelegramPlane } from "react-icons/fa";
import { useForm } from "react-hook-form";
import Datetime from "react-datetime";
import "react-datetime/css/react-datetime.css";
import moment from "moment";
import "moment/locale/ja";

type FormData = {
  message: string;
};

const ChatForm: VFC = () => {
  const [isReminded, setIsReminded] = useState(false);
  const { onOpen, onClose, isOpen } = useDisclosure();
  const [datetime, setDatetime] = useState<string>(moment().format());
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
        if (isReminded) {
          alert(
            JSON.stringify({
              type: 1,
              content: data.message,
              datetime: datetime,
            })
          );
        } else {
          alert(
            JSON.stringify({
              type: 2,
              content: data.message,
            })
          );
        }
        // もろもろリセットする
        onClose();
        setIsReminded(false);
        setDatetime(moment().format());
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
          <Popover
            isOpen={isOpen}
            onOpen={onOpen}
            onClose={onClose}
            placement="top"
            closeOnBlur={false}
          >
            <PopoverTrigger>
              <IconButton
                variant="ghost"
                aria-label="remind"
                icon={<BellIcon w={6} h={6} />}
                color={isReminded ? `${MAIN_COLOR}.400` : "gray.500"}
              />
            </PopoverTrigger>
            <PopoverContent p={5}>
              <PopoverArrow />
              <PopoverCloseButton />
              <PopoverHeader fontWeight="bold">
                メッセージのリマインド
              </PopoverHeader>
              <PopoverBody textAlign="center" py="4">
                <Datetime
                  initialValue={new Date()}
                  locale="ja"
                  onChange={(value) => {
                    if (typeof value === "string") console.log(value);
                    else setDatetime(value.format());
                  }}
                />
                <FormControl
                  display="flex"
                  align="center"
                  justifyContent="center"
                  w="full"
                  mt="3"
                >
                  <FormLabel htmlFor="remind-switch" mb="0">
                    にリマインドする
                  </FormLabel>
                  <Switch
                    id="remind-switch"
                    size="md"
                    isChecked={isReminded}
                    onChange={() => setIsReminded(!isReminded)}
                  />
                </FormControl>
              </PopoverBody>
            </PopoverContent>
          </Popover>
          <IconButton
            variant="ghost"
            aria-label="attach image"
            color="gray.500"
            icon={<AttachmentIcon w={6} h={6} />}
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
