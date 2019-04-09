import React, { PureComponent } from 'react';
import { connect } from 'dva';
import { Badge, Card, Form, Table } from 'antd';
import PageHeaderLayout from '../../layouts/PageHeaderLayout';
import MerchantCard from './MerchantCard';
import { formatDate } from '../../utils/utils';

import styles from './MerchantDetail.less';

@connect(state => ({
  loading: state.loading.models.merchant,
  merchant: state.merchant,
}))
@Form.create()
class MerchantDetail extends PureComponent {
  state = {
    selectedRowKeys: [],
  };

  componentDidMount() {
    this.dispatch({
      type: 'merchant/fetch',
      search: {},
      pagination: {},
    });
  }

  dispatch = action => {
    const { dispatch } = this.props;
    dispatch(action);
  };

  renderDataForm() {
    return <MerchantCard onCancel={this.onDataFormCancel} onSubmit={this.onDataFormSubmit} />;
  }

  render() {
    const {
      loading,
      merchant: {
        data: { list },
      },
    } = this.props;

    const { selectedRowKeys } = this.state;

    const columns = [
      {
        title: '编号',
        dataIndex: 'code',
      },
      {
        title: '名称',
        dataIndex: 'name',
      },
      {
        title: '地址',
        dataIndex: 'address',
      },
      {
        title: '地址描述',
        dataIndex: 'address_description',
      },
      {
        title: '备注',
        dataIndex: 'description',
      },
      {
        title: '状态',
        dataIndex: 'status',
        render: val => {
          if (val === 1) {
            return <Badge status="success" text="启用" />;
          }
          return <Badge status="error" text="停用" />;
        },
      },
      {
        title: '创建时间',
        dataIndex: 'created_at',
        render: val => <span>{formatDate(val, 'YYYY-MM-DD HH:mm')}</span>,
      },
    ];

    const breadcrumbList = [{ title: '商户管理' }, { title: '商户详情', href: '/merchant/detail' }];

    return (
      <PageHeaderLayout title="商户管理" breadcrumbList={breadcrumbList}>
        <Card bordered={false}>
          <div className={styles.tableList}>
            <div>
              <Table
                rowSelection={{
                  selectedRowKeys,
                  onChange: this.onTableSelectRow,
                }}
                loading={loading}
                rowKey={record => record.record_id}
                dataSource={list}
                columns={columns}
                onChange={this.onTableChange}
                size="small"
              />
            </div>
          </div>
        </Card>
        {this.renderDataForm()}
      </PageHeaderLayout>
    );
  }
}

export default MerchantDetail;
