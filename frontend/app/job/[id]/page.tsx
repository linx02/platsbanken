"use client";

import { getJob } from "@/app/fetchBackend";
import { useEffect, useState, Fragment } from "react";
import Link from "next/link";

export default function Page({ params }: { params: { id: string } }) {
  const [job, setJob] = useState<any>(null);
  const [lastApplicationDate, setLastApplicationDate] = useState<any>(null);

  useEffect(() => {
    const fetchJob = async () => {
      const job = await getJob(parseInt(params.id));
      job.description = job.description.replace(/\n/g, "<br />");
      setJob(job);
      setLastApplicationDate(formatDateAndDaysLeft(job.lastApplicationDate));
    };

    fetchJob();
  }, []);

  function renderDescription(description: string) {
    return description.split("\n").map((line, index) => (
      <Fragment key={index}>
        {line}
        <br />
      </Fragment>
    ));
  }

  function formatDateAndDaysLeft(dateString: string) {
    const monthNames = ["januari", "februari", "mars", "april", "maj", "juni", "juli", "augusti", "september", "oktober", "november", "december"];
    
    // Parse the input date
    const targetDate = new Date(dateString);
    const currentDate = new Date();

    // Format the day and month
    const day = targetDate.getUTCDate();
    const month = monthNames[targetDate.getUTCMonth()];

    // Calculate the days left
    const timeDiff = Number(targetDate) - Number(currentDate);
    const daysLeft = Math.ceil(timeDiff / (1000 * 60 * 60 * 24));

    // Return the formatted date and days left
    const data = {
        day: day,
        month: month,
        daysLeft: daysLeft
    }
    return data;
}

  return (
    <main className="bg-light-gray py-8">
      <div className="margin-x">
        <Link href="/" className="text-blue-800 font-medium text-lg flex items-center">
        <svg xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 24 24"><path fill="#263fa9" fill-rule="evenodd" d="m15 4l2 2l-6 6l6 6l-2 2l-8-8z"/></svg>
          Sökresultat
        </Link>
      </div>
      <div className="grid grid-cols-3 bg-white margin-x !my-8 p-8 gap-x-8">
      {job && (
        <div className="col-span-2">
          <img src={job.logotype} alt="Logo" className="max-h-12" />
          <h1 className="text-4xl pt-4 font-medium text-text-gray">
            {job.title}
          </h1>
          <p>{job.location}</p>
          <p className="text-2xl py-4 font-medium text-text-gray">
            {job.company.name}
          </p>
          <p className="text-xl font-medium text-text-gray">{job.occupation}</p>
          <p className="text-xl font-medium text-text-gray">
            Kommun: {job.workplace.municipality}
          </p>
          <p className="text-md pt-4 text-text-gray">
            Omfattning: {job.workTimeExtent}
          </p>
          <p className="text-md text-text-gray">Varaktighet: {job.duration}</p>
          <p className="text-md text-text-gray">
            Anställningsform: {job.employmentType}
          </p>
          <h2 className="text-2xl py-4 font-medium text-text-gray">
            Om jobbet
          </h2>
          <p dangerouslySetInnerHTML={{ __html: job.description }}></p>
          <h2 className="text-2xl py-4 font-medium text-text-gray">
            Om anställningen
          </h2>
          <h3 className="text-xl font-medium text-text-gray">Lön</h3>
          <p>{job.salaryDescription}</p>
          <p className="pt-2">
            <span className="font-medium text-text-gray">Lönetyp: </span>
            {job.salaryType}
          </p>
          { job.conditions &&
          <div>
            <h3 className="text-xl pt-4 font-medium text-text-gray">
            Anställningsvillkor
          </h3>
          <p>{renderDescription(job.conditions)}</p>
          </div>
          }
          <h3 className="text-xl pt-4 font-medium text-text-gray">
            Var ligger arbetsplatsen?
          </h3>
          {job.workplace.street ?
          <p>
            <div className="flex items-center">
                <svg xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 24 24"><path fill="#000056" d="M12 2C7.589 2 4 5.589 4 9.995C3.971 16.44 11.696 21.784 12 22c0 0 8.029-5.56 8-12c0-4.411-3.589-8-8-8m0 12c-2.21 0-4-1.79-4-4s1.79-4 4-4s4 1.79 4 4s-1.79 4-4 4"/></svg>
                {job.workplace.street}
            </div>
            <div className="ml-[1em]">
                {job.workplace.postCode} {job.workplace.city}
            </div>
          </p>
          : <p className="flex items-center space-x-1">
            <svg xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 24 24"><path fill="#000056" d="M12 2C7.589 2 4 5.589 4 9.995C3.971 16.44 11.696 21.784 12 22c0 0 8.029-5.56 8-12c0-4.411-3.589-8-8-8m0 12c-2.21 0-4-1.79-4-4s1.79-4 4-4s4 1.79 4 4s-1.79 4-4 4"/></svg>
            Arbetsplatsen ligger i kommunen 
            <span className="font-semibold text-text-gray">
              {job.workplace.municipality}
            </span> 
            <span>i</span>
            <span className="font-semibold text-text-gray">
              {job.workplace.region}
            </span>
          </p>}
          <h2 className="text-2xl py-4 font-medium text-text-gray">
            Arbetsgivaren
          </h2>
          <p>{job.company.name}</p>
          <a
            className="text-blue-800 font-semibold text-md flex items-center gap-1"
            href={job.company.webAddress}
            target="_blank"
          >
            <svg xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 20 20"><g fill="#263fa9"><path d="M10.707 10.707a1 1 0 0 1-1.414-1.414l6-6a1 1 0 1 1 1.414 1.414z"/><path d="M15 15v-3.5a1 1 0 1 1 2 0V16a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h4.5a1 1 0 0 1 0 2H5v10zm2-7a1 1 0 1 1-2 0V4a1 1 0 1 1 2 0z"/><path d="M12 5a1 1 0 1 1 0-2h4a1 1 0 1 1 0 2z"/></g></svg>
            {job.company.webAddress}
          </a>
          {job.company.streetAddress && 
          <div>
          <h3 className="text-xl pt-4 font-medium text-text-gray">
            Postadress
          </h3>
          <p>{job.company.name}</p>
          <p>{job.company.streetAddress}</p>
          <p>
            {job.company.postCode} {job.company.city}
          </p>
          </div>
          }
          <p className="pt-12">AnnonsId: {job.id}</p>
          <p>Publicerad: {job.publishedDate}</p>
        </div>
      )}
      <div className="col-span-1 mt-12 sticky top-0 h-fit">
        {job && <div className="bg-light-gray w-full border-l-[10px] border-dark-blue px-4 py-8">
            <h3 className="text-2xl font-semibold text-text-gray">Sök jobbet</h3>
            {lastApplicationDate && <p className="pt-2">Ansök senast <span className="font-semibold text-text-gray">{lastApplicationDate.day + ' ' + lastApplicationDate.month}</span> (om {lastApplicationDate.daysLeft} dagar)</p>}
            {job.application.reference && <p className="pt-2">Ange referens <span className="font-semibold text-text-gray">{job.application.reference}</span> i din ansökan</p>}
            {job.application.webAddress ?
            <div className="flex flex-col">
                <div className="flex items-center space-x-2">
                    <svg xmlns="http://www.w3.org/2000/svg" width="2em" height="2em" viewBox="0 0 16 16"><path fill="#000056" d="M0 8a8 8 0 1 1 16 0A8 8 0 0 1 0 8m7.5-6.923c-.67.204-1.335.82-1.887 1.855A8 8 0 0 0 5.145 4H7.5zM4.09 4a9.3 9.3 0 0 1 .64-1.539a7 7 0 0 1 .597-.933A7.03 7.03 0 0 0 2.255 4zm-.582 3.5c.03-.877.138-1.718.312-2.5H1.674a7 7 0 0 0-.656 2.5zM4.847 5a12.5 12.5 0 0 0-.338 2.5H7.5V5zM8.5 5v2.5h2.99a12.5 12.5 0 0 0-.337-2.5zM4.51 8.5a12.5 12.5 0 0 0 .337 2.5H7.5V8.5zm3.99 0V11h2.653c.187-.765.306-1.608.338-2.5zM5.145 12q.208.58.468 1.068c.552 1.035 1.218 1.65 1.887 1.855V12zm.182 2.472a7 7 0 0 1-.597-.933A9.3 9.3 0 0 1 4.09 12H2.255a7 7 0 0 0 3.072 2.472M3.82 11a13.7 13.7 0 0 1-.312-2.5h-2.49c.062.89.291 1.733.656 2.5zm6.853 3.472A7 7 0 0 0 13.745 12H11.91a9.3 9.3 0 0 1-.64 1.539a7 7 0 0 1-.597.933M8.5 12v2.923c.67-.204 1.335-.82 1.887-1.855q.26-.487.468-1.068zm3.68-1h2.146c.365-.767.594-1.61.656-2.5h-2.49a13.7 13.7 0 0 1-.312 2.5m2.802-3.5a7 7 0 0 0-.656-2.5H12.18c.174.782.282 1.623.312 2.5zM11.27 2.461c.247.464.462.98.64 1.539h1.835a7 7 0 0 0-3.072-2.472c.218.284.418.598.597.933M10.855 4a8 8 0 0 0-.468-1.068C9.835 1.897 9.17 1.282 8.5 1.077V4z"/></svg>
                    <h3 className="text-xl font-semibold text-text-gray pt-4">Ansök via arbetsgivarens webbplats</h3>
                </div>
                <a target="_blank" className="ml-[2em] mt-2 bg-dark-blue px-4 py-2 text-white w-fit rounded-lg font-medium underline hover:bg-[#1616ab] transition duration-300" href={job.application.webAddress}>Ansök här</a>
            </div>
             :
            <div className="flex flex-col">
                <div className="flex items-center space-x-2 pt-4">
                    <svg xmlns="http://www.w3.org/2000/svg" width="1.5em" height="1.5em" viewBox="0 0 24 24"><path fill="#000056" d="M22.596 20.093a11.435 11.435 0 0 1-4.48 2.943l-.08.025c-1.634.594-3.52.938-5.486.938h-.1h.005l-.189.001c-1.766 0-3.456-.327-5.012-.923l.096.032a11.476 11.476 0 0 1-3.93-2.474l.004.003a11.005 11.005 0 0 1-2.502-3.719l-.026-.073a12.447 12.447 0 0 1-.895-4.676l.001-.152v.008l-.001-.113c0-1.686.361-3.288 1.01-4.733l-.029.073a11.961 11.961 0 0 1 2.669-3.808l.004-.004A12.488 12.488 0 0 1 7.535.938l.084-.03C9.024.334 10.655.001 12.363.001h.098h-.005h.041c1.52 0 2.986.234 4.363.668l-.103-.028a11.174 11.174 0 0 1 3.724 1.951l-.023-.017a9.546 9.546 0 0 1 2.552 3.17l.025.055A9.935 9.935 0 0 1 24 10.364v-.012l.002.187a9.968 9.968 0 0 1-.559 3.301l.021-.07a7.526 7.526 0 0 1-1.443 2.494l.007-.009a5.856 5.856 0 0 1-2.015 1.485l-.037.015a5.973 5.973 0 0 1-2.409.498h-.018h.001a3.301 3.301 0 0 1-2.099-.617l.01.007a1.894 1.894 0 0 1-.782-1.534v-.015v.001h-.16a5.285 5.285 0 0 1-1.48 1.456l-.02.012a4.256 4.256 0 0 1-2.487.697h.007a4.182 4.182 0 0 1-3.418-1.447l-.005-.005a5.676 5.676 0 0 1-1.205-3.79v.013c.001-.96.167-1.88.473-2.735l-.018.057a7.648 7.648 0 0 1 1.312-2.37l-.011.014a6.69 6.69 0 0 1 1.985-1.643l.035-.017a5.369 5.369 0 0 1 2.612-.626h-.004a3.96 3.96 0 0 1 2.109.525l-.019-.01c.503.281.89.717 1.103 1.242l.006.017h.032l.262-1.291h2.903l-1.28 6.101c-.043.3-.102.632-.177 1.002c-.069.31-.11.666-.113 1.032v.002l-.001.063c0 .314.08.61.22.868l-.005-.01a.87.87 0 0 0 .834.369h-.004a2.439 2.439 0 0 0 2.099-1.341l.006-.014a6.858 6.858 0 0 0 .829-3.659l.001.016l.001-.155c0-1.183-.241-2.31-.676-3.335l.021.056a7.078 7.078 0 0 0-1.801-2.513l-.007-.006a7.747 7.747 0 0 0-2.698-1.517l-.055-.015a11.087 11.087 0 0 0-3.378-.515l-.115.001h.006l-.122-.001a9.55 9.55 0 0 0-3.707.744l.063-.024A8.63 8.63 0 0 0 5.715 5.34l-.003.003a9.081 9.081 0 0 0-1.854 2.938l-.021.062a10.524 10.524 0 0 0-.675 3.747c0 1.372.257 2.684.727 3.89l-.025-.073a8.406 8.406 0 0 0 1.968 2.917l.003.003a8.597 8.597 0 0 0 2.972 1.835l.06.019c1.114.406 2.4.641 3.741.641l.144-.001h-.007a11.131 11.131 0 0 0 4.62-.865l-.074.027a12.562 12.562 0 0 0 3.513-2.331l-.006.006zM12.391 8.356h-.038c-.504 0-.968.169-1.338.455l.005-.004a3.938 3.938 0 0 0-.987 1.128l-.01.019a5.782 5.782 0 0 0-.613 1.491l-.009.041a6.486 6.486 0 0 0-.214 1.609v.004c.002.293.031.578.085.854l-.005-.029c.057.306.171.578.331.817l-.005-.008a1.84 1.84 0 0 0 1.686.851h-.006l.078.001a2.53 2.53 0 0 0 1.431-.44l-.009.006a3.522 3.522 0 0 0 1.004-1.082l.009-.016c.248-.415.446-.895.567-1.405l.007-.035c.106-.421.171-.907.179-1.407v-.005c0-.36-.035-.712-.102-1.052l.006.034a2.7 2.7 0 0 0-.35-.917l.007.012a1.932 1.932 0 0 0-.661-.654l-.009-.005a1.926 1.926 0 0 0-.977-.262l-.066.001h.003z"/></svg>
                    <h3 className="text-xl font-semibold text-text-gray">Ansök via mejl</h3>
                </div>
            <p className="ml-[2em] mt-2">Mejla din ansökan till</p>
            <a className="ml-[2em] text-blue-800 underline mt-0" href={"mailto:" + job.application.email}>{job.application.email}</a>
        </div>
            }
        </div>}
      </div>
    </div>
    </main>
  );
}
