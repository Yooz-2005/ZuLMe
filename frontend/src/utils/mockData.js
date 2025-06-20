// 模拟预订数据
export const mockReservations = [
  {
    id: 'RES001',
    vehicle: {
      id: 1,
      brand: 'Range Rover',
      name: 'Evoque',
      images: [
        'https://images.unsplash.com/photo-1549399736-8e3c8b8b8b8b?w=400',
        'https://images.unsplash.com/photo-1549399736-8e3c8b8b8b8b?w=400'
      ]
    },
    start_date: '2024-01-15',
    end_date: '2024-01-18',
    pickup_location: '北京首都国际机场',
    return_location: '北京首都国际机场',
    total_amount: 2400,
    status: 'pending_payment',
    created_at: '2024-01-10 14:30:00'
  },
  {
    id: 'RES002',
    vehicle: {
      id: 2,
      brand: 'Bentley',
      name: 'Continental GT',
      images: [
        'https://images.unsplash.com/photo-1549399736-8e3c8b8b8b8b?w=400'
      ]
    },
    start_date: '2024-01-20',
    end_date: '2024-01-22',
    pickup_location: '上海虹桥机场',
    return_location: '上海虹桥机场',
    total_amount: 3600,
    status: 'confirmed',
    created_at: '2024-01-12 10:15:00'
  },
  {
    id: 'RES003',
    vehicle: {
      id: 3,
      brand: 'Mercedes-Benz',
      name: 'S-Class',
      images: [
        'https://images.unsplash.com/photo-1549399736-8e3c8b8b8b8b?w=400'
      ]
    },
    start_date: '2024-01-08',
    end_date: '2024-01-10',
    pickup_location: '广州白云机场',
    return_location: '广州白云机场',
    total_amount: 1800,
    status: 'completed',
    created_at: '2024-01-05 16:45:00'
  },
  {
    id: 'RES004',
    vehicle: {
      id: 4,
      brand: 'BMW',
      name: 'X7',
      images: [
        'https://images.unsplash.com/photo-1549399736-8e3c8b8b8b8b?w=400'
      ]
    },
    start_date: '2024-01-25',
    end_date: '2024-01-28',
    pickup_location: '深圳宝安机场',
    return_location: '深圳宝安机场',
    total_amount: 2700,
    status: 'processing',
    created_at: '2024-01-14 09:20:00'
  },
  {
    id: 'RES005',
    vehicle: {
      id: 5,
      brand: 'Audi',
      name: 'A8L',
      images: [
        'https://images.unsplash.com/photo-1549399736-8e3c8b8b8b8b?w=400'
      ]
    },
    start_date: '2024-01-12',
    end_date: '2024-01-14',
    pickup_location: '杭州萧山机场',
    return_location: '杭州萧山机场',
    total_amount: 1500,
    status: 'cancelled',
    created_at: '2024-01-08 11:30:00'
  }
];

// 模拟订单数据
export const mockOrders = [
  {
    id: 'ORD001',
    reservation_id: 'RES002',
    vehicle: {
      id: 2,
      brand: 'Bentley',
      name: 'Continental GT'
    },
    total_amount: 3600,
    status: 'paid',
    payment_status: 'completed',
    created_at: '2024-01-12 10:20:00'
  }
];

// 模拟支付数据
export const mockPayments = [
  {
    id: 'PAY001',
    order_id: 'ORD001',
    amount: 3600,
    payment_method: 'alipay',
    status: 'completed',
    payment_url: 'https://example.com/alipay/pay?order=ORD001',
    created_at: '2024-01-12 10:25:00'
  }
];
