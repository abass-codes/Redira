export interface DashboardSummary {
    TotalLinks: number;
    TotalClicks: number;
    ActiveLinks: number;
  }
  
  
  export interface TimelinePoint {
    Date: string;
    Clicks: number;
  }
  
  
  export interface LinkAnalytics {
    ClickedAt: string;
    IpAddress: string;
    UserAgent: string;
    Referer: string | null;
  }