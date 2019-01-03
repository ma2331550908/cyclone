import { Layout, Menu, Icon } from 'antd';
import React, { Component } from 'react';
import { NavLink } from 'react-router-dom';
import Routes from '../routes';
import { withRouter } from 'react-router-dom';
import PropTypes from 'prop-types';
import BreadCrumb from './Breadcrumb';

const { Header, Content, Sider } = Layout;

class CoreLayout extends Component {
  static propTypes = {
    location: PropTypes.shape({
      pathname: PropTypes.string,
    }),
  };

  state = {
    collapsed: false,
  };

  onCollapse = collapsed => {
    this.setState({ collapsed });
  };

  render() {
    const { location } = this.props;
    const pathSnippets = location.pathname.split('/').filter(i => i);
    const selectNav = pathSnippets[0] ? pathSnippets[0] : '/overview';
    return (
      <Layout style={{ height: '100%' }}>
        <Header className="header" style={{ lineHeight: '64px' }}>
          <div className="logo">CYCLONE</div>
          <Menu mode="horizontal" theme="dark">
            <Menu.Item key="3">
              <Icon type="user" />
            </Menu.Item>
          </Menu>
        </Header>
        <Layout>
          <Sider
            width={200}
            collapsible
            style={{ background: '#fff' }}
            collapsed={this.state.collapsed}
            onCollapse={this.onCollapse}
          >
            <Menu
              mode="inline"
              defaultSelectedKeys={[selectNav]}
              style={{ height: '100%', borderRight: 0 }}
            >
              <Menu.Item key="overview">
                <NavLink to="/overview" activeClassName="active">
                  <Icon type="home" />
                  <span>{intl.get('sideNav.overview')}</span>
                </NavLink>
              </Menu.Item>
              <Menu.Item key="project">
                <NavLink to="/project" activeClassName="active">
                  <Icon type="project" />
                  <span>{intl.get('sideNav.project')}</span>
                </NavLink>
              </Menu.Item>
              <Menu.Item key="resource">
                <NavLink to="/resource" activeClassName="active">
                  <Icon type="cluster" />
                  <span>{intl.get('sideNav.resource')}</span>
                </NavLink>
              </Menu.Item>
              <Menu.Item key="template">
                <NavLink to="/template" activeClassName="active">
                  <Icon type="profile" />
                  <span>{intl.get('sideNav.template')}</span>
                </NavLink>
              </Menu.Item>
              <Menu.Item key="workflow">
                <NavLink to="/workflow" activeClassName="active">
                  <Icon type="share-alt" />
                  <span>{intl.get('sideNav.workflow')}</span>
                </NavLink>
              </Menu.Item>
              <Menu.Item key="integration">
                <NavLink to="/integration" activeClassName="active">
                  <Icon type="sliders" />
                  <span>{intl.get('sideNav.integration')}</span>
                </NavLink>
              </Menu.Item>
              {/* TODO: manage and setting */}
              {/* <SubMenu
                key="manage"
                title={
                  <span>
                    <Icon type="team" />
                    管理中心
                  </span>
                }
              >
                <Menu.Item key="tenant">租户</Menu.Item>
                <Menu.Item key="user">
                  <NavLink activeClassName="active">
                    用户
                  </NavLink>
                </Menu.Item>
              </SubMenu>
              <SubMenu
                key="setting"
                title={
                  <span>
                    <Icon type="setting" />
                    配置中心
                  </span>
                }
              >
                <Menu.Item>
                  <span>设置</span>
                </Menu.Item>
              </SubMenu> */}
            </Menu>
          </Sider>
          <Layout style={{ margin: '16px' }}>
            <BreadCrumb location={location} />
            <Content
              style={{
                background: '#fff',
                padding: 24,
                minHeight: 280,
              }}
            >
              <Routes />
            </Content>
          </Layout>
        </Layout>
      </Layout>
    );
  }
}

export default withRouter(CoreLayout);