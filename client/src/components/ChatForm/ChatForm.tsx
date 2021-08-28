import { VFC, useState, useCallback } from "react";
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
  Image,
  Container,
} from "@chakra-ui/react";
import { MAIN_COLOR } from "../../constants/MainColor";
import { AttachmentIcon, BellIcon, CloseIcon } from "@chakra-ui/icons";
import { FaTelegramPlane } from "react-icons/fa";
import { useForm } from "react-hook-form";
import Datetime from "react-datetime";
import "react-datetime/css/react-datetime.css";
import moment from "moment";
import "moment/locale/ja";
import { useDropzone } from "react-dropzone";
import axios from "axios";
import { User } from "../../type/User";

type FormData = {
  message: string;
};

type Props = {
  recruit_id: number;
  currentUser: User;
};

const ChatForm: VFC<Props> = ({ recruit_id, currentUser }) => {
  const [isReminded, setIsReminded] = useState(false);
  const [src, setSrc] = useState("");
  const [myFiles, setMyFiles] = useState<File[]>([]);
  const { onOpen, onClose, isOpen } = useDisclosure();
  const [datetime, setDatetime] = useState<string>(moment().format());
  const {
    handleSubmit,
    register,
    setValue,
    formState: { errors, isSubmitting },
  } = useForm<FormData>();
  const onDrop = useCallback(async (acceptedFiles: File[]) => {
    if (!acceptedFiles[0]) return;
    try {
      setMyFiles([...acceptedFiles]);
      handlePreview(acceptedFiles);
    } catch (error) {
      alert(error);
    }
  }, []);
  const onDropRejected = () => {
    alert("画像のみ受け付けることができます。");
  };
  // プレビューを表示
  const handlePreview = (files: any) => {
    if (files === null) {
      return;
    }
    const file = files[0];
    if (file === null) {
      return;
    }
    var reader = new FileReader();
    reader.readAsDataURL(file);
    reader.onload = () => {
      setSrc(reader.result as string);
    };
  };
  const { getRootProps, getInputProps } = useDropzone({
    onDrop,
    onDropRejected,
    accept: "image/*",
  });

  const onSubmit = handleSubmit(async (data) => {
    if (myFiles[0]) {
      const data = new FormData();
      data.append("image", myFiles[0]);
      data.append("type", "3");
      const headers = {
        "content-type": "multipart/form-data",
        Authorization: `Bearer ${currentUser.token}`,
      };
      // TODO: Check
      const res = await axios.post("/message/" + recruit_id, data, { headers });
      console.log(res);
    } else if (isReminded) {
      if (data.message.length === 0) return;
      // TODO: Check
      axios
        .post("/message/" + recruit_id, {
          type: 1,
          content: data.message,
          execute_at: datetime,
          headers: { Authorization: `Bearer ${currentUser.token}` },
        })
        .then((res) => console.log(res))
        .catch((e) => console.log(e));
    } else {
      if (data.message.length === 0) return;
      // TODO: Check
      axios
        .post("/message/" + recruit_id, {
          type: 2,
          content: data.message,
          headers: { Authorization: `Bearer ${currentUser.token}` },
        })
        .then((res) => console.log(res))
        .catch((e) => console.log(e));
    }
    // もろもろリセットする
    onClose();
    setIsReminded(false);
    setDatetime(moment().format());
    setValue("message", "");
    setMyFiles([]);
    setSrc("");
  });

  return (
    <form onSubmit={onSubmit}>
      <FormControl
        id="inputText2"
        isInvalid={errors.message ? true : false}
        textAlign="center"
      >
        {src && (
          <Container pos="absolute" top="-300px" justify="center">
            <Image w="auto" h="300px" src={src} />
            <IconButton
              pos="absolute"
              left="0"
              top="0"
              aria-label="delete image"
              onClick={() => {
                setMyFiles([]);
                setSrc("");
              }}
              icon={<CloseIcon w={6} h={6} />}
            />
          </Container>
        )}
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

          <div {...getRootProps()}>
            <input {...getInputProps()} />
            <IconButton
              variant="ghost"
              aria-label="attach image"
              color="gray.500"
              icon={<AttachmentIcon w={6} h={6} />}
            />
          </div>
          <Input
            variant="filled"
            placeholder="テキストを入力してください"
            aria-label="input message"
            ml="2"
            {...register("message", {
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
