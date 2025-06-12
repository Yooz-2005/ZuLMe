import React, { useState } from 'react';
import { Form, Input, Select, Button, DatePicker, Typography, Space, Checkbox, message, Row, Col } from 'antd';
import moment from 'moment';
import instance from '../../../utils/axiosConfig'; // 導入配置好的 axios 實例，修正路徑

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
    // 您可以在此處添加更多省份和城市數據
};

const MyInfoPage = () => {
    const [form] = Form.useForm();
    const [selectedProvince, setSelectedProvince] = useState(null);
    const [selectedCity, setSelectedCity] = useState(null);

    console.log('手機號碼從 localStorage 獲取:', localStorage.getItem('phoneNumber'));

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
        // 在這裡調用後端接口更新用户信息
        // alert('信息已保存 (模擬)'); // Remove mock alert

        const token = localStorage.getItem('token'); // 這裡仍然保留獲取 token 的邏輯，但不再手動添加到 headers
        console.log('從 localStorage 獲取的 token:', token);

        if (!token) {
            message.error('未檢測到登入憑證，請重新登入。');
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
            emailSubscription: values.emailSubscription,
        };
        console.log('發送到後端的 payload:', payload); // 添加 payload 日誌

        try {
            const response = await instance.post('/user/profile', payload); // 使用配置好的 instance

            if (response.data.code === 200) {
                message.success('個人信息保存成功！');
            } else {
                message.error(response.data.msg || '保存失敗，請重試。');
            }
        } catch (error) {
            console.error('保存個人信息時出錯:', error);
            message.error(error.response?.data?.msg || '保存失敗，請稍後重試。');
        }
    };

    return (
        <Form
            form={form}
            layout="horizontal"
            onFinish={onFinish}
            labelCol={{ span: 6 }}
            wrapperCol={{ span: 16 }}
            initialValues={{
                emailSubscription: false,
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
                    <Text>{localStorage.getItem('phoneNumber') ? localStorage.getItem('phoneNumber').replace(/(\d{3})\d{4}(\d{4})/, '$1****$2') : ''}</Text>
                    <Button type="link">修改</Button>
                </Space>
            </Form.Item>

            <Form.Item label="电子邮箱" name="email">
                <Space style={{ width: '100%' }}>
                    <Input placeholder="请输入常用邮箱" style={{ flex: 1 }} />
                    <Button>验证</Button>
                </Space>
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
                    {selectedProvince && (
                        <Form.Item name="city" noStyle>
                            <Select
                                placeholder="请选择市"
                                style={{ width: 120 }}
                                onChange={handleCityChange}
                            >
                                {citiesData[selectedProvince] && Object.keys(citiesData[selectedProvince]).map(city => (
                                    <Option key={city} value={city}>{city}</Option>
                                ))}
                            </Select>
                        </Form.Item>
                    )}
                    {selectedCity && (
                        <Form.Item name="district" noStyle>
                            <Input placeholder="无需重复写省市" style={{ flex: 1 }} />
                        </Form.Item>
                    )}
                </Space>
            </Form.Item>

            <Form.Item label="紧急联系人" name="emergencyName">
                <Input placeholder="请输入紧急联系人姓名" />
            </Form.Item>

            <Form.Item label="联系电话" name="emergencyPhone">
                <Input placeholder="请输入紧急联系人电话" />
            </Form.Item>

            <Form.Item label="邮件订阅" name="emailSubscription" valuePropName="checked">
                <Checkbox>
                    为您推送租借优惠活动信息
                </Checkbox>
                {form.getFieldValue('emailSubscription') && (
                    <Text type="success" style={{ marginLeft: '16px' }}>已订阅</Text>
                )}
            </Form.Item>

            <Form.Item wrapperCol={{ offset: 6, span: 16 }}>
                <Button type="primary" htmlType="submit">
                    保存
                </Button>
            </Form.Item>
        </Form>
    );
};

export default MyInfoPage; 