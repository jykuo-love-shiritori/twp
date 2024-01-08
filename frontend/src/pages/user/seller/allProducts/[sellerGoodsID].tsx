import '@style/global.css';

import { Col, Row } from 'react-bootstrap';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faFileUpload, faPen, faTrash } from '@fortawesome/free-solid-svg-icons';
import { useRef, useState } from 'react';
import { useMutation } from '@tanstack/react-query';
import { SubmitHandler, useForm } from 'react-hook-form';
import { useNavigate, useParams } from 'react-router-dom';
import React, { useReducer } from 'react';

import TButton from '@components/TButton';
import FormItem from '@components/FormItem';
import { GetUserName, useAuth } from '@lib/Auth';

import {
  TagProps,
  ProductProps,
  RequestProps,
  tagStyle,
  LeftBgStyle,
  isDataValid,
  SetFormData,
  TagsAction,
  DeleteTagsAction,
  GoodsImgStyle,
  GoodsBGStyle,
} from './NewGoods';

export interface TagPropsAnotherVersion {
  name: string;
  tag_id: number;
}
export interface GetResponseProps {
  product_info: {
    description: string;
    enabled: boolean;
    expire_date: string;
    id: number;
    image_url: string;
    name: string;
    price: number;
    sales: number;
    stock: number;
  };
  tags: TagPropsAnotherVersion[];
}

interface PatchResponseProps {
  id: number;
  name: string;
  description: string;
  price: number;
  image_url: string;
  expire_date: string;
  edit_date: string;
  stock: number;
  sales: number;
  enabled: boolean;
}

interface ConnectionTagsPros {
  id?: number;
  tag_id: number;
}

const ADD_TAG = 'ADD_TAG';
const DELETE_TAG = 'DELETE_TAG';
const SET_TAG = 'SET_TAG';

const tagsReducer = (state: TagProps[], action: TagsAction) => {
  switch (action.type) {
    case ADD_TAG:
      return [...state, action.payload];
    case DELETE_TAG:
      return state.filter((_, i) => i !== action.payload);
    case SET_TAG:
      return action.payload;
    default:
      return state;
  }
};

const deleteTagsAction = (index: number): DeleteTagsAction => {
  return {
    type: DELETE_TAG,
    payload: index,
  };
};

const EachSellerGoods = () => {
  const token = useAuth();
  const navigate = useNavigate();
  const username = GetUserName();

  const params = useParams<{ goods_id?: string }>();
  let goods_id: number | undefined;

  if (params.goods_id) {
    goods_id = parseInt(params.goods_id);
  }

  const [tag, setTag] = useState('');
  const [reducerTags, dispatchTags] = useReducer(tagsReducer, []);
  const [queryTags, setQueryTags] = useState<string[]>([]);
  const [file, setFile] = useState<string | null>(null);

  const { register, setValue, handleSubmit } = useForm<ProductProps>({
    defaultValues: {
      name: 'new product',
      description: 'new product description',
      price: 0,
      image: undefined,
      expire_date: 'expire date',
      stock: 0,
      enabled: true,
      tags: [],
    },
  });

  const addTag = useMutation({
    mutationFn: async (data: RequestProps) => {
      const response = await fetch('/api/seller/tag', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          Accept: 'application/json',
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify(data),
      });
      if (!response.ok) {
        throw new Error('add tag failed');
      }
      return await response.json();
    },

    onSuccess: async (responseData: TagProps) => {
      dispatchTags({ type: ADD_TAG, payload: responseData });
      updateTags();
      await connectTag.mutate({ id: goods_id, tag_id: responseData.id });
    },
  });

  const queryTag = useMutation({
    mutationFn: async (data: string) => {
      const response = await fetch(`/api/seller/tag?name=${data}`, {
        headers: {
          Accept: 'application/json',
          Authorization: `Bearer ${token}`,
        },
      });
      if (!response.ok) {
        throw new Error('query tag failed');
      }
      return await response.json();
    },
    onSuccess: (responseData: TagProps[]) => {
      const tagNames = responseData.map((tag) => tag.name);
      setQueryTags(tagNames);
    },
  });

  const queryProduct = useMutation({
    mutationFn: async (id: number) => {
      const response = await fetch(`/api/seller/product/${id}`, {
        headers: {
          Accept: 'application/json',
          Authorization: `Bearer ${token}`,
        },
      });
      if (!response.ok) {
        throw new Error('query tag failed');
      }
      return await response.json();
    },
    onSuccess: (responseData: GetResponseProps) => {
      setValue('name', responseData.product_info.name);
      setValue('description', responseData.product_info.description);
      setValue('price', responseData.product_info.price);
      setFile(responseData.product_info.image_url);
      setValue('image', responseData.product_info.image_url);
      setValue(
        'expire_date',
        new Date(responseData.product_info.expire_date).toLocaleDateString('en-CA'),
      );
      setValue('stock', responseData.product_info.stock);
      setValue('enabled', Boolean(responseData.product_info.enabled));
      const convertedTags = responseData.tags.map((tag) => ({
        id: tag.tag_id,
        name: tag.name,
      }));
      dispatchTags({ type: SET_TAG, payload: convertedTags });
      setValue(
        'tags',
        reducerTags.map((tag) => tag.id),
      );

      goods_id = responseData.product_info.id;
    },
  });

  const updateProduct = useMutation({
    mutationFn: async (data: ProductProps) => {
      if (!isDataValid(data)) {
        throw new Error('Invalid data');
      }

      const formData = SetFormData(data);

      const response = await fetch(`/api/seller/product/${goods_id}`, {
        method: 'PATCH',
        headers: {
          Accept: 'application/json',
          Authorization: `Bearer ${token}`,
        },
        body: formData,
        redirect: 'follow',
      });
      if (!response.ok) {
        throw new Error('add product failed');
      }
      return await response.json();
    },
    onSuccess: (responseData: PatchResponseProps) => {
      setValue('name', responseData.name);
      setValue('description', responseData.description);
      setValue('price', responseData.price);
      setFile(responseData.image_url);
      setValue('image', responseData.image_url);
      setValue('expire_date', new Date(responseData.expire_date).toLocaleDateString('en-CA'));
      setValue('stock', responseData.stock);
      setValue('enabled', Boolean(responseData.enabled));
      navigate('/user/seller/manageProducts');
    },
  });

  const connectTag = useMutation({
    mutationFn: async (data: ConnectionTagsPros) => {
      if (goods_id === undefined) {
        throw new Error('Invalid goods_id');
      }

      const response = await fetch(`/api/seller/product/${goods_id}/tag`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          Accept: 'application/json',
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify(data),
      });
      if (!response.ok) {
        throw new Error('connect tag failed');
      }
      return await response.json();
    },
  });

  const removeConnectedTag = useMutation({
    mutationFn: async (data: ConnectionTagsPros) => {
      if (goods_id === undefined) {
        throw new Error('Invalid goods_id');
      }

      const response = await fetch(`/api/seller/product/${goods_id}/tag`, {
        method: 'DELETE',
        headers: {
          'Content-Type': 'application/json',
          Accept: 'application/json',
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify(data),
      });
      if (!response.ok) {
        throw new Error('remove connected tag failed');
      }
      return await response.json();
    },
  });

  const deleteProduct = useMutation({
    mutationFn: async () => {
      const response = await fetch(`/api/seller/product/${goods_id}`, {
        method: 'DELETE',
        headers: {
          'Content-Type': 'application/json',
          Accept: 'application/json',
          Authorization: `Bearer ${token}`,
        },
      });
      if (!response.ok) {
        throw new Error('delete product failed');
      }
      return await response.json();
    },
    onSuccess: () => {
      navigate('/user/seller/manageProducts');
    },
  });

  const addNewTag = async (event: React.KeyboardEvent<HTMLInputElement>) => {
    if (event.keyCode === 229) return;

    if (event.key === 'Enter') {
      const input = event.currentTarget.value.trim();

      if (input == '') {
        Reset();
        return;
      }

      if (reducerTags.some((currentTag) => currentTag.name === tag)) {
        Reset();
        return;
      }

      queryTag.mutate(tag, {
        onSuccess: async (responseData) => {
          const existingTag = responseData.find((currentTag) => currentTag.name === tag);

          if (existingTag) {
            dispatchTags({ type: ADD_TAG, payload: existingTag });
            updateTags();
            await connectTag.mutate({ id: goods_id, tag_id: existingTag.id });
          } else {
            if (username === undefined) {
              return;
            }
            addTag.mutate({ name: tag, seller_name: username });
          }

          Reset();
        },
      });
    }
  };

  const TagOnChange = async (e: React.ChangeEvent<HTMLInputElement>) => {
    const tagName = e.target.value;
    setQueryTags([]);

    if (tagName == '') {
      Reset();
      return;
    }
    setTag(tagName);
    queryTag.mutate(tagName);
  };

  const deleteTag = async (index: number) => {
    await removeConnectedTag.mutate({ id: goods_id, tag_id: reducerTags[index].id });
    dispatchTags(deleteTagsAction(index));
  };

  const updateTags = () => {
    setValue(
      'tags',
      reducerTags.map((tag) => tag.id),
    );
  };

  const Reset = () => {
    setTag('');
    setQueryTags([]);
  };

  const uploadFile = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files) {
      if (!e.target.files[0].name.match(/\.(jpg|jpeg|png|gif)$/i)) {
        alert('not an image');
        return;
      } else {
        setValue('image', e.target.files[0]);
        setFile(URL.createObjectURL(e.target.files[0]));
      }
    }
  };

  const prevGoodsIdRef = useRef<number | null>(null);

  if (goods_id && goods_id !== prevGoodsIdRef.current) {
    queryProduct.mutate(goods_id);
    prevGoodsIdRef.current = goods_id;
  }

  const onSubmit: SubmitHandler<ProductProps> = async (data) => {
    updateProduct.mutate(data);
  };

  return (
    <div style={{ padding: '55px 12% 0 12%' }}>
      <form onSubmit={handleSubmit(onSubmit)}>
        <Row>
          <Col xs={12} md={5} style={LeftBgStyle}>
            <div className='flex_wrapper' style={{ padding: '0 8% 10% 8%' }}>
              <div
                style={{
                  position: 'relative',
                  width: '100%',
                  height: '100%',
                  borderRadius: '0 0 30px 0',
                }}
              >
                <div>
                  {file ? (
                    <div style={{ overflow: ' hidden' }}>
                      <img src={file} alt='File preview' style={GoodsImgStyle} />
                    </div>
                  ) : (
                    <div style={GoodsBGStyle}>
                      <div style={{ padding: '30% 5% 30% 5%' }} className='center'>
                        <FontAwesomeIcon icon={faFileUpload} size='6x' />
                      </div>
                    </div>
                  )}
                </div>
                <br />
                <Row
                  style={{
                    position: 'absolute',
                    zIndex: '1',
                    right: '0px',
                    bottom: '40px',
                  }}
                >
                  <Col xs={9}></Col>
                  <Col xs={3}>
                    <form method='post' encType='multipart/form-data'>
                      <label
                        htmlFor='file'
                        className='custom-file-upload'
                        style={{ minWidth: '40px' }}
                      >
                        <div className='button pointer center' style={{ padding: '10px' }}>
                          <FontAwesomeIcon icon={faPen} className='white_word' />
                        </div>
                      </label>

                      <input
                        id='file'
                        name='file'
                        type='file'
                        style={{ display: 'none' }}
                        onChange={uploadFile}
                      />
                    </form>
                  </Col>
                </Row>
              </div>
              <br />
              <span className='dark'>add more tags</span>

              <input
                type='text'
                placeholder='Input tags'
                className='quantity_box'
                value={tag}
                onChange={TagOnChange}
                onKeyDown={addNewTag}
                style={{ marginBottom: '10px' }}
                list='queryTags'
              />
              <datalist id='queryTags'>
                {queryTags.map((tag, index) => (
                  <option key={index} value={tag} />
                ))}
              </datalist>

              <Row xs='auto'>
                {reducerTags.map((currentTag, index) => (
                  <Col style={tagStyle} key={index} className='center'>
                    <Row style={{ width: '100%' }} className='center'>
                      <Col xs={1} className='center'>
                        <FontAwesomeIcon
                          icon={faTrash}
                          className='white_word pointer'
                          onClick={() => deleteTag(index)}
                        />
                      </Col>
                      <Col xs={9} lg={10}>
                        <span style={{ wordBreak: 'break-all' }}>{currentTag.name}</span>
                      </Col>
                    </Row>
                  </Col>
                ))}
              </Row>

              <div style={{ height: '50px' }} />
              <TButton
                text='Delete Product'
                action={() => goods_id !== undefined && deleteProduct.mutate()}
              />
              <TButton text='Confirm Changes' action={handleSubmit(onSubmit)} />
            </div>
          </Col>
          <Col xs={12} md={7}>
            <div style={{ padding: '7% 0% 7% 0%' }}>
              <FormItem label='Product Name'>
                <input type='text' {...register('name', { required: true })} />
              </FormItem>

              <FormItem label='Product Price'>
                <input type='number' {...register('price', { required: true })} />
              </FormItem>

              <FormItem label='Product Quantity'>
                <input type='number' {...register('stock', { required: true })} />
              </FormItem>

              <FormItem label='Product Introduction'>
                <textarea {...register('description', { required: true })} />
              </FormItem>

              <FormItem label='Best Before Date'>
                <input type='date' {...register('expire_date', { required: true })} />
              </FormItem>

              <FormItem label='Product Listing' isCenter={false}>
                <input type='checkbox' {...register('enabled')} />
              </FormItem>
            </div>
          </Col>
        </Row>
      </form>
    </div>
  );
};

export default EachSellerGoods;
