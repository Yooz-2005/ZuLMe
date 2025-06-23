import React, { useState } from 'react';
import { Form, Input, Select, Button, DatePicker, Typography, Space, message, Modal } from 'antd';
import moment from 'moment';
import instance from '../../../utils/axiosConfig'; // 导入配置好的 axios 实例，修正路径

const { Option } = Select;
const { Title, Text } = Typography;

const citiesData = {
    '江苏省': {
        '南京市': ['玄武区', '秦淮区', '鼓楼区'],
        '苏州市': ['姑苏区', '吴中区', '相城区'],
        '泰州市': ['海陵区', '高港区', '姜堰区'],
    },
    '浙江省': {
        '杭州市': ['西湖区', '上城区', '拱墅区'],
        '宁波市': ['海曙区', '江北区', '鄞州区'],
    },
    // 您可以在此处添加更多省份和城市数据
};

const MyInfoPage = ({ onPhoneUpdate }) => {
    const [form] = Form.useForm();
    const [updatePhoneForm] = Form.useForm(); // 新增用于手机号码修改的表单实例
    const [selectedProvince, setSelectedProvince] = useState(null);
    const [selectedCity, setSelectedCity] = useState(null);
    const [isPhoneModalVisible, setIsPhoneModalVisible] = useState(false); // 控制手机号码修改模态框的可见性

    console.log('手机号码从 localStorage 获取:', localStorage.getItem('userPhone'));

    const handleProvinceChange = (value) => {
        setSelectedProvince(value);
        setSelectedCity(null); // Reset city when province changes
        form.setFieldsValue({ city: null, district: null }); // Clear city and district fields
    };

    const handleCityChange = (value) => {
        setSelectedCity(value);
        form.setFieldsValue({ district: null }); // Clear district when city changes
    };

    // 移除模擬數據加載，用戶將自行填寫
    /*
    React.useEffect(() => {
        // 假设从后端获取到的用户信息
        const userInfo = {
            realName: '丁雨欣',
            idType: '1',
            idNumber: '321284200502078028',
            idExpireDate: '2033-02-18',
            phoneNumber: '178****7929',
            email: '3106373537@qq.com',
            province: '江苏省',
            city: '泰州市',
            district: '姜堰区',
            emergencyName: '某某某',
            emergencyPhone: '138XXXXXXXX',
            emailSubscription: true, // 假设邮件订阅状态为 true，以便显示"已订阅"
        };

        form.setFieldsValue({
            realName: userInfo.realName,
            idType: userInfo.idType,
            idNumber: userInfo.idNumber,
            idExpireDate: userInfo.idExpireDate ? moment(userInfo.idExpireDate) : null,
            email: userInfo.email,
            province: userInfo.province,
            city: userInfo.city,
            district: userInfo.district,
            emergencyName: userInfo.emergencyName,
            emergencyPhone: userInfo.emergencyPhone,
            emailSubscription: userInfo.emailSubscription,
        });
    }, [form]);
    */

    const onFinish = async (values) => {
        console.log('Received values of form:', values);
        // 在这里调用后端接口更新用户信息
        // alert('信息已保存 (模拟)'); // Remove mock alert

        const token = localStorage.getItem('token'); // 这里仍然保留获取 token 的逻辑，但不再手动添加到 headers
        console.log('从 localStorage 获取的 token:', token);

        if (!token) {
            message.error('未检测到登录凭证，请重新登录。');
            return;
        }

        const payload = {
            real_name: values.realName,
            id_type: values.idType,
            id_number: values.idNumber,
            id_expire_date: values.idExpireDate ? values.idExpireDate.format('YYYY-MM-DD') : null, // 格式化日期
            email: values.email,
            province: values.province,
            city: values.city,
            district: values.district,
            emergency_name: values.emergencyName,
            emergency_phone: values.emergencyPhone,
        };
        console.log('发送到后端的 payload:', payload); // 添加 payload 日志

        try {
            const response = await instance.post('/user/profile', payload); // 使用配置好的 instance

            if (response.data.code === 200) {
                message.success('个人信息保存成功！');
            } else {
                message.error(response.data.msg || '保存失败，请重试。');
            }
        } catch (error) {
            console.error('保存个人信息时出错:', error);
            message.error(error.response?.data?.msg || '保存失败，请稍后重试。');
        }
    };

    // 显示手机号码修改模态框
    const showPhoneModal = () => {
        setIsPhoneModalVisible(true);
    };

    // 隐藏手机号码修改模态框
    const handlePhoneModalCancel = () => {
        setIsPhoneModalVisible(false);
        updatePhoneForm.resetFields(); // 清空表单字段
    };

    // 发送手机验证码
    const sendPhoneCode = async () => {
        try {
            const newPhoneNumber = updatePhoneForm.getFieldValue('newPhoneNumber');
            if (!newPhoneNumber) {
                message.error('请输入手机号码');
                return;
            }
            // 调用发送验证码的后端接口
            const response = await instance.post('/user/sendCode', { phone: newPhoneNumber, source: "updatePhone" });
            if (response.data.code === 200) {
                message.success('验证码已发送');
            } else {
                message.error(response.data.msg || '验证码发送失败');
            }
        } catch (error) {
            console.error('发送验证码时出错:', error);
            message.error(error.response?.data?.msg || '验证码发送失败，请稍后重试。');
        }
    };

    // 提交手机号码修改
    const onFinishUpdatePhone = async (values) => {
        console.log('Received values for phone update:', values);
        try {
            const token = localStorage.getItem('token');
            if (!token) {
                message.error('未检测到登录凭证，请重新登录。');
                return;
            }
            // 调用修改手机号码的后端接口
            const response = await instance.post('/user/phone', {
                phone: values.newPhoneNumber,
                code: values.verificationCode,
            });

            if (response.data.code === 200) {
                message.success('手机号码修改成功！');
                localStorage.setItem('userPhone', values.newPhoneNumber); // 更新本地存储的手机号码
                if (onPhoneUpdate) {
                    onPhoneUpdate(values.newPhoneNumber); // 通知父组件手机号码已更新
                }
                setIsPhoneModalVisible(false); // 隐藏模态框
                updatePhoneForm.resetFields(); // 清空表单字段
            } else {
                message.error(response.data.msg || '手机号码修改失败，请重试。');
            }
        } catch (error) {
            console.error('修改手机号码时出错:', error);
            message.error(error.response?.data?.msg || '修改手机号码失败，请稍后重试。');
        }
    };

    return (
        <>
            <Form
                form={form}
                layout="horizontal"
                onFinish={onFinish}
                labelCol={{ span: 6 }}
                wrapperCol={{ span: 16 }}
                initialValues={{
                    realName: null,
                    idType: null,
                    idNumber: null,
                    idExpireDate: null,
                    email: null,
                    province: null,
                    city: null,
                    district: null,
                    emergencyName: null,
                    emergencyPhone: null,
                }}
                style={{ maxWidth: '800px', margin: '0 auto', padding: '24px', background: '#fff', borderRadius: '8px' }}
            >
                <Title level={4} style={{ marginBottom: '24px', textAlign: 'center' }}>我的信息</Title>

                <Form.Item label="姓名" name="realName">
                    <Input placeholder="请输入真实姓名" />
                </Form.Item>

                <Form.Item label="证件">
                    <Space>
                        <Form.Item name="idType" noStyle>
                            <Select placeholder="请选择证件类型" style={{ width: 180 }}>
                                <Option value="1">身份证</Option>
                                <Option value="2">台湾居民来往大陆通行证</Option>
                                <Option value="3">港澳居民来往大陆通行证</Option>
                                <Option value="4">外籍护照</Option>
                            </Select>
                        </Form.Item>
                        <Form.Item name="idNumber" noStyle>
                            <Input placeholder="请输入证件号码" style={{ flex: 1 }} />
                        </Form.Item>
                    </Space>
                </Form.Item>

                <Form.Item label="有效期" name="idExpireDate">
                    <DatePicker placeholder="请选择有效期" style={{ width: '100%' }} />
                </Form.Item>

                <Form.Item label="手机号码">
                    <Space>
                        <Text>{localStorage.getItem('userPhone') ? localStorage.getItem('userPhone').replace(/(\d{3})\d{4}(\d{4})/, '$1****$2') : ''}</Text>
                        <Button type="link" onClick={showPhoneModal}>修改</Button>
                    </Space>
                </Form.Item>

                <Form.Item label="电子邮箱" name="email">
                    <Input placeholder="请输入常用邮箱" style={{ width: '100%' }} />
                </Form.Item>

                <Form.Item label="通讯地址">
                    <Space style={{ width: '100%' }}>
                        <Form.Item name="province" noStyle>
                            <Select
                                placeholder="请选择省"
                                style={{ width: 120 }}
                                onChange={handleProvinceChange}
                            >
                                {Object.keys(citiesData).map(province => (
                                    <Option key={province} value={province}>{province}</Option>
                                ))}
                            </Select>
                        </Form.Item>
                        <Form.Item name="city" noStyle>
                            <Select
                                placeholder="请选择市"
                                style={{ width: 120 }}
                                onChange={handleCityChange}
                                disabled={!selectedProvince}
                            >
                                {selectedProvince && citiesData[selectedProvince] &&
                                    Object.keys(citiesData[selectedProvince]).map(city => (
                                        <Option key={city} value={city}>{city}</Option>
                                    ))}
                            </Select>
                        </Form.Item>
                        <Form.Item name="district" noStyle>
                            <Select
                                placeholder="请选择区"
                                style={{ width: 120 }}
                                disabled={!selectedCity}
                            >
                                {selectedCity && citiesData[selectedProvince][selectedCity] &&
                                    citiesData[selectedProvince][selectedCity].map(district => (
                                        <Option key={district} value={district}>{district}</Option>
                                    ))}
                            </Select>
                        </Form.Item>
                    </Space>
                </Form.Item>

                <Form.Item label="紧急联系人" name="emergencyName">
                    <Input placeholder="请输入紧急联系人姓名" />
                </Form.Item>

                <Form.Item label="紧急联系电话" name="emergencyPhone">
                    <Input placeholder="请输入紧急联系人电话" />
                </Form.Item>

                <Form.Item wrapperCol={{ offset: 6, span: 16 }}>
                    <Button type="primary" htmlType="submit">
                        保存
                    </Button>
                </Form.Item>
            </Form>

            <Modal
                title="修改手机号码"
                open={isPhoneModalVisible}
                onCancel={handlePhoneModalCancel}
                footer={null}
            >
                <Form
                    form={updatePhoneForm}
                    onFinish={onFinishUpdatePhone}
                    layout="vertical"
                >
                    <Form.Item
                        name="newPhoneNumber"
                        label="新手机号码"
                        rules={[
                            { required: true, message: '请输入新手机号码' },
                            { pattern: /^1[3-9]\d{9}$/, message: '请输入正确的手机号码' }
                        ]}
                    >
                        <Input placeholder="请输入新手机号码" />
                    </Form.Item>

                    <Form.Item
                        name="verificationCode"
                        label="验证码"
                        rules={[
                            { required: true, message: '请输入验证码' },
                            { pattern: /^\d{5}$/, message: '验证码必须是5位数字' }
                        ]}
                    >
                        <Input placeholder="请输入验证码" />
                    </Form.Item>

                    <Form.Item>
                        <Button type="primary" htmlType="submit" style={{ marginRight: '8px' }}>
                            确认修改
                        </Button>
                        <Button onClick={handlePhoneModalCancel}>
                            取消
                        </Button>
                    </Form.Item>
                </Form>
            </Modal>
        </>
    );
};

export default MyInfoPage;