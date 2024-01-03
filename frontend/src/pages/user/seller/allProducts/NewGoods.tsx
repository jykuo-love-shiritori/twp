import '@style/global.css';

import { Col, Row } from 'react-bootstrap';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faFileUpload, faPen, faTrash } from '@fortawesome/free-solid-svg-icons';
import { useState } from 'react';
import { useMutation } from '@tanstack/react-query';
import { SubmitHandler, useForm } from 'react-hook-form';
import { useNavigate } from 'react-router-dom';
import React, { useReducer, useEffect, CSSProperties } from 'react';

import TButton from '@components/TButton';
import FormItem from '@components/FormItem';
import { useAuth } from '@lib/Auth';

export const GoodsImgStyle: CSSProperties = {
  borderRadius: '0 0 30px 0',
  width: '100%',
  maxHeight: '40vh',
  minHeight: '40vh',
  objectFit: 'cover',
};

export const GoodsBGStyle: CSSProperties = {
  borderRadius: '0 0 30px 0',
  backgroundColor: 'black',
  width: '100%',
  maxHeight: '40vh',
  minHeight: '40vh',
};

export interface TagProps {
  id: number;
  name: string;
}

export interface ProductProps {
  name: string;
  description: string;
  price: number;
  image: File | string;
  expire_date: string;
  stock: number;
  enable: boolean;
  tags: number[];
}

export interface RequestProps {
  name: string;
  seller_name: string;
}

// eslint-disable-next-line react-refresh/only-export-components
export const tagStyle = {
  borderRadius: '30px',
  background: ' var(--button_light)',
  padding: '1% 1% 1% 3%',
  color: 'white',
  margin: '5px 0 5px 5px',
  width: '100%',
};

export const LeftBgStyle = {
  backgroundColor: 'rgba(255, 255, 255, 0.7)',
  boxShadow: '6px 4px 10px 2px rgba(0, 0, 0, 0.25)',
};

// i want to alert for all exception, letting user know what happened
// eslint-disable-next-line react-refresh/only-export-components
export const isDataValid = (data: ProductProps) => {
  if (!/^-?\d+$/.test(data.price.toString())) {
    alert("can't input texts or point numbers on price");
    return false;
  }

  if (!/^-?\d+$/.test(data.stock.toString())) {
    alert("can't input texts or point numbers on stock");
    return false;
  }

  if (data.price <= 0) {
    alert("price can't be 0 or smaller than 0!");
    return false;
  }

  if (data.stock < 0) {
    alert("stock can't be smaller than 0!");
    return false;
  }

  if (data.image == undefined) {
    alert('must has product image!');
    return false;
  }

  if (new Date() > new Date(data.expire_date)) {
    alert('product already expired!');
    return false;
  }

  if (data.tags == undefined) {
    alert('must have one tag');
    return false;
  }

  return true;
};

export const SetFormData = (data: ProductProps) => {
  const formData = new FormData();
  formData.append('name', data.name);
  formData.append('description', data.description);
  formData.append('image', data.image);
  formData.append('price', String(data.price));
  formData.append('expire_date', new Date(data.expire_date).toISOString());
  formData.append('stock', String(data.stock));
  formData.append('enable', String(data.enable));
  formData.append('tags', data.tags.join(','));
  return formData;
};

const ADD_TAG = 'ADD_TAG';
const DELETE_TAG = 'DELETE_TAG';

export type TagsAction = AddTagsAction | DeleteTagsAction;

export interface AddTagsAction {
  type: typeof ADD_TAG;
  payload: TagProps;
}

export interface DeleteTagsAction {
  type: typeof DELETE_TAG;
  payload: number;
}

const tagsReducer = (state: TagProps[], action: TagsAction) => {
  switch (action.type) {
    case ADD_TAG:
      return [...state, action.payload];
    case DELETE_TAG:
      return state.filter((_, i) => i !== action.payload);
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

const EmptyGoods = () => {
  const token = useAuth();
  const navigate = useNavigate();

  const [tag, setTag] = useState('');
  const [tags, setTags] = useState<TagProps[]>([]);
  const [queryTags, setQueryTags] = useState<string[]>([]);
  const [file, setFile] = useState<string | null>(null);

  const [reducerTags, dispatchTags] = useReducer(tagsReducer, []);

  useEffect(() => {
    setTags(reducerTags);
  }, [reducerTags, tags]);

  const { register, setValue, handleSubmit } = useForm<ProductProps>({
    defaultValues: {
      name: 'new product',
      description: 'new product description',
      price: 0,
      image: undefined,
      expire_date: 'expire date',
      stock: 0,
      enable: true,
      tags: undefined,
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

    onSuccess: (responseData: TagProps) => {
      dispatchTags({ type: ADD_TAG, payload: responseData });
      updateTags(responseData);
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

  const addProduct = useMutation({
    mutationFn: async (data: ProductProps) => {
      if (!isDataValid(data)) {
        throw new Error('Invalid data');
      }

      data.tags = tags.map((tag) => tag.id);
      const formData = SetFormData(data);

      const response = await fetch('/api/seller/product', {
        method: 'POST',
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

      if (tags.some((currentTag) => currentTag.name === tag)) {
        Reset();
        return;
      }

      queryTag.mutate(tag, {
        onSuccess: (responseData) => {
          const existingTag = responseData.find((currentTag) => currentTag.name === tag);

          if (existingTag) {
            dispatchTags({ type: ADD_TAG, payload: existingTag });
            updateTags(existingTag);
          } else {
            // TODO : seller name need to be change to corresponding user
            addTag.mutate({ name: tag, seller_name: 'user1' });
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

  const deleteTag = (index: number) => {
    dispatchTags(deleteTagsAction(index));
    setTags((prevTags) => prevTags.filter((_, i) => i !== index));
  };

  const updateTags = (tag: TagProps) => {
    setTags((prevTags) => {
      const newTags = [...prevTags, tag];
      setValue(
        'tags',
        newTags.map((tag) => tag.id),
      );
      return newTags;
    });
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

  const onSubmit: SubmitHandler<ProductProps> = async (data) => {
    await addProduct.mutate(data);
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
                {tags.map((currentTag, index) => (
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
                action={() => navigate('/user/seller/manageProducts')}
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
                <input type='number' {...register('stock', { required: true, min: 0 })} />
              </FormItem>

              <FormItem label='Product Introduction'>
                <textarea {...register('description', { required: true })} />
              </FormItem>

              <FormItem label='Best Before Date'>
                <input
                  type='date'
                  {...register('expire_date', { required: true, validate: { date: () => true } })}
                />
              </FormItem>
            </div>
          </Col>
        </Row>
      </form>
    </div>
  );
};

export default EmptyGoods;
